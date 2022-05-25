package svc

import (
	"github.com/zeromicro/go-zero/zrpc"
	"mall/service/user/api/internal/config"
	"mall/service/user/rpc/user"
)

type ServiceContext struct {
	Config config.Config

	//RPC User Service
	UserRPC user.USer
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:  c,
		UserRPC: user.NewUSer(zrpc.MustNewClient(c.UserRPC)),
	}
}
