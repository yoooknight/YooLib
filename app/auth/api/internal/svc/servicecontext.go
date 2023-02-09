package svc

import (
	"app/auth/api/internal/config"
	"app/middlewares"
	"app/tools/mysql"
	"gorm.io/gorm"
	"net/http"
)

type ServiceContext struct {
	Config         config.Config
	AuthMiddleware func(next http.HandlerFunc) http.HandlerFunc
	Mysql          *gorm.DB
}

func NewServiceContext(c config.Config) *ServiceContext {
	var ctx = &ServiceContext{
		Config: c,
		Mysql:  mysql.GetDB(c.Mysql),
	}

	ctx.AuthMiddleware = middlewares.NewAuthMiddleware(c.Auth.AccessSecret, c.SuperManager.Username).Handle
	return ctx
}
