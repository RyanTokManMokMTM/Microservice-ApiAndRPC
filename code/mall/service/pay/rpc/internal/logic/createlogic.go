package logic

import (
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/grpc/status"
	"mall/service/order/rpc/types/order"
	"mall/service/pay/model"
	"mall/service/pay/rpc/internal/svc"
	"mall/service/pay/rpc/types/pay"
	"mall/service/user/rpc/user"
)

type CreateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateLogic {
	return &CreateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateLogic) Create(in *pay.CreateRequest) (*pay.CreateResponse, error) {
	// todo: add your logic here and delete this line

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

	//is paid
	_, err = l.svcCtx.PayModel.FindOneByOid(l.ctx, in.Oid)
	if err == nil {
		return nil, status.Error(100, "order has been paid")
	}

	newPay := model.Pay{
		Uid:    in.Uid,
		Oid:    in.Oid,
		Amount: in.Amount,
		Source: 0, //?
		Status: 0,
	}

	res, err := l.svcCtx.PayModel.Insert(l.ctx, &newPay)
	if err != nil {
		return nil, status.Error(500, err.Error())
	}

	newPay.Id, err = res.LastInsertId()
	if err != nil {
		return nil, status.Error(500, err.Error())
	}

	return &pay.CreateResponse{
		Id: newPay.Id,
	}, nil
}
