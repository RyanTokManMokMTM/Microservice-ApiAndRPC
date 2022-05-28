package svc

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/zrpc"
	"mall/service/order/model"
	"mall/service/order/rpc/internal/config"
	"mall/service/product/rpc/product"
	"mall/service/user/rpc/user"
)

type ServiceContext struct {
	Config     config.Config
	OrderModel model.OrderModel

	UserRPC    user.USer
	ProductRPC product.Product
}

func NewServiceContext(c config.Config) *ServiceContext {
	sqlConn := sqlx.NewMysql(c.Mysql.DataSource)
	return &ServiceContext{
		Config:     c,
		OrderModel: model.NewOrderModel(sqlConn, c.CacheRedis),
		UserRPC:    user.NewUSer(zrpc.MustNewClient(c.UserPRC)),
		ProductRPC: product.NewProduct(zrpc.MustNewClient(c.ProductRPC)),
	}
}
