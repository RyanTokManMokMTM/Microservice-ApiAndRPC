package logic

import (
	"context"
	"fmt"
	"google.golang.org/grpc/status"
	"mall/service/order/model"

	"mall/service/order/rpc/internal/svc"
	"mall/service/order/rpc/types/order"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteLogic {
	return &DeleteLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteLogic) Delete(in *order.DeleteRequest) (*order.DeleteResponse, error) {
	// todo: add your logic here and delete this line
	//Find order id
	_, err := l.svcCtx.OrderModel.FindOne(l.ctx, in.Id)
	if err != nil {
		if err == model.ErrNotFound {
			return nil, status.Error(100, fmt.Sprintf("Order not exists"))
		}
		return nil, status.Error(500, err.Error())
	}

	err = l.svcCtx.OrderModel.Delete(l.ctx, in.Id)
	if err != nil {
		return nil, status.Error(500, err.Error())
	}

	return &order.DeleteResponse{}, nil
}
