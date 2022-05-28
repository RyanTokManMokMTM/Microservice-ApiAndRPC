package svc

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/zrpc"
	"mall/service/order/rpc/order"
	"mall/service/pay/model"
	"mall/service/pay/rpc/internal/config"
	"mall/service/user/rpc/user"
)

type ServiceContext struct {
	Config   config.Config
	PayModel model.PayModel

	UserRPC  user.USer
	OrderRPC order.Order
}

func NewServiceContext(c config.Config) *ServiceContext {
	sqlConn := sqlx.NewMysql(c.Mysql.DataSource)
	return &ServiceContext{
		Config:   c,
		PayModel: model.NewPayModel(sqlConn, c.CacheRedis),
		UserRPC:  user.NewUSer(zrpc.MustNewClient(c.UserRPC)),
		OrderRPC: order.NewOrder(zrpc.MustNewClient(c.OrderRPC)),
	}
}
