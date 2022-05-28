package logic

import (
	"context"
	"fmt"
	"google.golang.org/grpc/status"
	"mall/service/user/model"
	"mall/service/user/rpc/type/user"

	"mall/service/order/rpc/internal/svc"
	"mall/service/order/rpc/types/order"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListLogic {
	return &ListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ListLogic) List(in *order.ListRequest) (*order.ListResponse, error) {
	// todo: add your logic here and delete this line
	_, err := l.svcCtx.UserRPC.UserInfo(l.ctx, &user.UserInfoRequest{
		Id: in.Uid,
	})
	if err != nil {
		return nil, err
	}

	//getting all orders by uid
	res, err := l.svcCtx.OrderModel.FindAllOrder(l.ctx, in.Uid)
	if err != nil {
		if err == model.ErrNotFound {
			return nil, status.Error(100, fmt.Sprintf("user not exists"))
		}
		return nil, status.Error(500, err.Error())
	}

	orderLine := make([]*order.DetailResponse, 0)
	for _, item := range res {
		orderLine = append(orderLine, &order.DetailResponse{
			Id:     item.Id,
			Uid:    item.Uid,
			Pid:    item.Pid,
			Amount: item.Amount,
			Status: item.Status,
		})
	}

	return &order.ListResponse{
		Data: orderLine,
	}, nil
}
