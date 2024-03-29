package logic

import (
	"context"
	"fmt"
	"google.golang.org/grpc/status"
	"mall/service/user/model"

	"mall/service/product/rpc/internal/svc"
	"mall/service/product/rpc/types/product"

	"github.com/zeromicro/go-zero/core/logx"
)

type DetailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DetailLogic {
	return &DetailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DetailLogic) Detail(in *product.DetailRequest) (*product.DetailResponse, error) {
	// todo: add your logic here and delete this line
	result, err := l.svcCtx.ProductModel.FindOne(l.ctx, in.Id)
	if err != nil {
		if err == model.ErrNotFound {
			return nil, status.Error(100, fmt.Sprintf("Product not exists"))
		}
		return nil, status.Error(500, err.Error())
	}
	return &product.DetailResponse{
		Id:     result.Id,
		Name:   result.Name,
		Desc:   result.Desc,
		Stock:  result.Stock,
		Amount: result.Amount,
		Status: result.Status,
	}, nil
}
