package logic

import (
	"context"

	"mall/service/user/api/internal/svc"
	"mall/service/user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CheckHealthLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCheckHealthLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CheckHealthLogic {
	return &CheckHealthLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CheckHealthLogic) CheckHealth() (resp *types.PongResponse, err error) {
	// todo: add your logic here and delete this line

	return &types.PongResponse{
		Message: "pong",
	}, nil
}
