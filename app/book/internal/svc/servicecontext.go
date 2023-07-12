package svc

import (
	"app/auth/rpc/authrpc"
	"app/book/internal/config"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config  config.Config
	AuthRpc authrpc.AuthRPC
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:  c,
		AuthRpc: authrpc.NewAuthRPC(zrpc.MustNewClient(c.AuthRpc)),
	}
}
