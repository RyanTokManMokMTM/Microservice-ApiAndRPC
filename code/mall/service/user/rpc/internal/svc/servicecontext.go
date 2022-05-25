package svc

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"mall/service/user/model"
	"mall/service/user/rpc/internal/config"
)

type ServiceContext struct {
	Config    config.Config
	UserModel model.UserModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	sqlConn := sqlx.NewMysql(c.Mysql.DataSource) //mysql conn
	return &ServiceContext{
		Config:    c,
		UserModel: model.NewUserModel(sqlConn, c.CacheRedis),
	}
}
