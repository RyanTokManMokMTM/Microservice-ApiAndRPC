package logic

import (
	"context"
	"mall/service/order/rpc/types/order"

	"mall/service/order/api/internal/svc"
	"mall/service/order/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListLogic {
	return &ListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ListLogic) List(req *types.ListRequest) (resp []*types.ListResponse, err error) {
	// todo: add your logic here and delete this line
	res, err := l.svcCtx.OrderRPC.List(l.ctx, &order.ListRequest{
		Uid: req.Uid,
	})
	if err != nil {
		return nil, err
	}

	orders := make([]*types.ListResponse, 0)
	for _, order := range res.Data {
		orders = append(orders, &types.ListResponse{
			Id:     order.Id,
			Uid:    order.Uid,
			Pid:    order.Pid,
			Amount: order.Amount,
			Status: order.Status,
		})
	}
	return orders, nil
}
