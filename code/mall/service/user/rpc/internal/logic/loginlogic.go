package logic

import (
	"context"
	"fmt"
	"google.golang.org/grpc/status"
	"mall/common/cryptx"
	"mall/service/user/model"

	"mall/service/user/rpc/internal/svc"
	"mall/service/user/rpc/type/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *LoginLogic) Login(in *user.LoginRequest) (*user.LoginResponse, error) {
	// todo: add your logic here and delete this line

	//check user
	res, err := l.svcCtx.UserModel.FindOneByMobile(l.ctx, in.Mobile)
	if err != nil {
		if err == model.ErrNotFound {
			return nil, status.Error(100, fmt.Sprintf("user not exist"))
		}
		return nil, status.Error(500, err.Error())

	}

	//checking password
	passwords := cryptx.PasswordEncrypt(l.svcCtx.Config.Salt, in.Password)
	if passwords != res.Password {
		return nil, status.Error(100, fmt.Sprintf("Passowrd incorrect"))
	}
	//TO Internal RPC/API Only
	return &user.LoginResponse{
		Id:     res.Id,
		Name:   res.Name,
		Gender: res.Gender,
		Mobile: res.Mobile,
	}, nil
}
