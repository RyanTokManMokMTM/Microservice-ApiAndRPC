package logic

import (
	"context"
	"fmt"
	"google.golang.org/grpc/status"
	"mall/service/user/model"

	"mall/service/user/rpc/internal/svc"
	"mall/service/user/rpc/type/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserInfoLogic {
	return &UserInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UserInfoLogic) UserInfo(in *user.UserInfoRequest) (*user.UserInfoResponse, error) {
	// todo: add your logic here and delete this line

	//RPC Get User Info
	//check user exists
	res, err := l.svcCtx.UserModel.FindOne(l.ctx, in.Id) //select by
	if err != nil {
		if err == model.ErrNotFound {
			return nil, status.Error(100, fmt.Sprintf("user not exists"))
		}
		return nil, status.Error(500, err.Error())
	}

	return &user.UserInfoResponse{
		Id:     res.Id,
		Gender: res.Gender,
		Mobile: res.Mobile,
		Name:   res.Name,
	}, nil
}
