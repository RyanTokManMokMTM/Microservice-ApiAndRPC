package logic

import (
	"context"
	"mall/common/jwt"
	"mall/service/user/rpc/type/user"
	"time"

	"mall/service/user/api/internal/svc"
	"mall/service/user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req *types.LoginRequest) (resp *types.LoginResponse, err error) {
	// todo: add your logic here and delete this line
	//login API
	//Calling RPC

	res, err := l.svcCtx.UserRPC.Login(l.ctx, &user.LoginRequest{
		Mobile:   req.Mobile,
		Password: req.Password,
	})

	if err != nil {
		return nil, err
	}

	timeNow := time.Now().Unix()
	expired := l.svcCtx.Config.Auth.AccessExpire
	token, err := jwt.GetToken(l.svcCtx.Config.Auth.AccessSecret, timeNow, expired, res.Id)
	if err != nil {
		return nil, err
	}
	return &types.LoginResponse{
		AccessToken:   token,
		AccessExpired: timeNow + expired,
	}, nil
}
