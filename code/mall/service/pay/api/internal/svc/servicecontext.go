package svc

import (
	"github.com/zeromicro/go-zero/zrpc"
	"mall/service/pay/api/internal/config"
	"mall/service/pay/rpc/pay"
)

type ServiceContext struct {
	Config config.Config
	PayRPC pay.Pay
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		PayRPC: pay.NewPay(zrpc.MustNewClient(c.PayRPC)),
	}
}
