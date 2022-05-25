package logic

import (
	"context"
	"google.golang.org/grpc/status"
	"mall/common/cryptx"
	"mall/service/user/model"

	"mall/service/user/rpc/internal/svc"
	"mall/service/user/rpc/type/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RegisterLogic) Register(in *user.RegisterRequest) (*user.RegisterResponse, error) {
	// todo: add your logic here and delete this line

	//check mobile exists or not
	_, err := l.svcCtx.UserModel.FindOneByMobile(l.ctx, in.Mobile)

	if err == model.ErrNotFound {
		//create a new user
		m := model.User{
			Name:     in.Name,
			Mobile:   in.Mobile,
			Gender:   in.Gender,
			Password: cryptx.PasswordEncrypt(l.svcCtx.Config.Salt, in.Password),
		}

		result, err := l.svcCtx.UserModel.Insert(l.ctx, &m)
		if err != nil {
			return nil, status.Error(500, err.Error())
		}

		m.Id, err = result.LastInsertId()
		if err != nil {
			return nil, status.Error(500, err.Error())
		}

		return &user.RegisterResponse{
			Id:     m.Id,
			Name:   m.Name,
			Mobile: m.Mobile,
			Gender: m.Gender,
		}, nil
	}

	return nil, status.Error(500, err.Error())
}
