package logic

import (
	"context"
	"google.golang.org/grpc/status"
	"mall/service/order/rpc/types/order"
	"mall/service/user/model"
	"mall/service/user/rpc/type/user"

	"mall/service/pay/rpc/internal/svc"
	"mall/service/pay/rpc/types/pay"

	"github.com/zeromicro/go-zero/core/logx"
)

type CallbackLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCallbackLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CallbackLogic {
	return &CallbackLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CallbackLogic) Callback(in *pay.CallbackRequest) (*pay.CallbackResponse, error) {
	// todo: add your logic here and delete this line
	//pay!!!

	//check user
	_, err := l.svcCtx.UserRPC.UserInfo(l.ctx, &user.UserInfoRequest{Id: in.Uid})
	if err != nil {
		return nil, err
	}

	//check order
	_, err = l.svcCtx.OrderRPC.Detail(l.ctx, &order.DetailRequest{Id: in.Oid})
	if err != nil {
		return nil, err
	}

	//check paid
	res, err := l.svcCtx.PayModel.FindOne(l.ctx, in.Id)
	if err != nil {
		if err == model.ErrNotFound {
			return nil, status.Error(100, "paid record not exist")
		}
		return nil, status.Error(500, err.Error())
	}

	//compare record and input
	if in.Amount != res.Amount {
		return nil, status.Error(100, "no enough amount to pay")
	}

	res.Source = in.Source //paid method
	res.Status = in.Status //update state

	err = l.svcCtx.PayModel.Update(l.ctx, res)
	if err != nil {
		return nil, status.Error(500, err.Error())
	}

	//Update Order paid status
	//Update state to 1
	_, err = l.svcCtx.OrderRPC.Paid(l.ctx, &order.PaidRequest{
		Id: res.Oid,
	})
	if err != nil {
		return nil, status.Error(500, err.Error())
	}
	return &pay.CallbackResponse{}, nil
}
