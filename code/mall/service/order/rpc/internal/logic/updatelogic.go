package logic

import (
	"context"
	"fmt"
	"google.golang.org/grpc/status"
	"mall/service/user/model"

	"mall/service/order/rpc/internal/svc"
	"mall/service/order/rpc/types/order"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateLogic {
	return &UpdateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateLogic) Update(in *order.UpdateRequest) (*order.UpdateResponse, error) {
	// todo: add your logic here and delete this line
	//find order
	res, err := l.svcCtx.OrderModel.FindOne(l.ctx, in.Id)
	if err != nil {
		if err == model.ErrNotFound {
			return nil, status.Error(100, fmt.Sprintf("Order not exist"))
		}
		return nil, status.Error(500, err.Error())
	}

	if in.Uid != 0 {
		res.Uid = in.Uid
	}

	if in.Pid != 0 {
		res.Pid = in.Pid
	}

	if in.Amount != 0 {
		res.Amount = in.Amount
	}

	if in.Status != 0 {
		res.Status = in.Status
	}

	//update our order
	err = l.svcCtx.OrderModel.Update(l.ctx, res)
	if err != nil {
		return nil, status.Error(500, err.Error())
	}

	return &order.UpdateResponse{}, nil
}
