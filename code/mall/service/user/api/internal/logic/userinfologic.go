package logic

import (
	"context"
	"encoding/json"
	"log"
	"mall/service/user/rpc/type/user"

	"mall/service/user/api/internal/svc"
	"mall/service/user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserInfoLogic {
	return &UserInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserInfoLogic) UserInfo() (resp *types.UserInfoResponse, err error) {
	// todo: add your logic here and delete this line
	//calling RPC
	//from JWT
	uid, _ := l.ctx.Value("uid").(json.Number).Int64()
	log.Println(l.ctx)
	res, err := l.svcCtx.UserRPC.UserInfo(l.ctx, &user.UserInfoRequest{
		Id: uid,
	})

	if err != nil {
		return nil, err

	}

	return &types.UserInfoResponse{
		Id:     res.Id,
		Name:   res.Name,
		Gender: res.Gender,
		Mobile: res.Mobile,
	}, nil
}
