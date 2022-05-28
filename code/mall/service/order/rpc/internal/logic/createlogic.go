package logic

import (
	"context"
	"fmt"
	"google.golang.org/grpc/status"
	"mall/service/order/model"
	"mall/service/product/rpc/types/product"
	"mall/service/user/rpc/type/user"

	"mall/service/order/rpc/internal/svc"
	"mall/service/order/rpc/types/order"

	"github.com/zeromicro/go-zero/core/logx"
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

func (l *CreateLogic) Create(in *order.CreateRequest) (*order.CreateResponse, error) {
	// todo: add your logic here and delete this line

	//Check user
	_, err := l.svcCtx.UserRPC.UserInfo(l.ctx, &user.UserInfoRequest{
		Id: in.Uid,
	})

	if err != nil {
		return nil, err
	}

	//Check product
	productResult, err := l.svcCtx.ProductRPC.Detail(l.ctx, &product.DetailRequest{
		Id: in.Pid,
	})

	if err != nil {
		return nil, err
	}

	//check stock
	if productResult.Stock <= 0 {
		return nil, status.Error(500, fmt.Sprintf("Product is not enough"))
	}

	//Create an order
	newOrder := model.Order{
		Uid:    in.Uid,
		Pid:    in.Pid,
		Amount: in.Amount,
		Status: 0,
	}

	res, err := l.svcCtx.OrderModel.Insert(l.ctx, &newOrder)
	if err != nil {
		return nil, status.Error(500, err.Error())
	}

	newOrder.Id, err = res.LastInsertId()
	if err != nil {
		return nil, status.Error(500, err.Error())
	}

	//update product stock , -1
	_, err = l.svcCtx.ProductRPC.Update(l.ctx, &product.UpdateRequest{
		Id:     productResult.Id,
		Name:   productResult.Name,
		Desc:   productResult.Desc,
		Amount: productResult.Amount,
		Stock:  productResult.Stock - 1,
		Status: productResult.Status,
	})

	if err != nil {
		return nil, status.Error(500, err.Error())
	}
	return &order.CreateResponse{
		Id: newOrder.Id,
	}, nil
}
