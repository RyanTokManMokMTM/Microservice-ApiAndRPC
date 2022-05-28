package logic

import (
	"context"
	"mall/service/pay/rpc/types/pay"

	"mall/service/pay/api/internal/svc"
	"mall/service/pay/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DetailLogic {
	return &DetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DetailLogic) Detail(req *types.DetailReqeust) (resp *types.DetailResponse, err error) {
	// todo: add your logic here and delete this line
	res, err := l.svcCtx.PayRPC.Detail(l.ctx, &pay.DetailRequest{
		Id: req.Id,
	})
	if err != nil {
		return nil, err
	}

	return &types.DetailResponse{
		Id:     res.Id,
		Uid:    res.Uid,
		Oid:    res.Uid,
		Amount: res.Amount,
		Status: res.Status,
	}, nil
}
