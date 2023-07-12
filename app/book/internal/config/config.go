package config

import (
	"app/tools/mysql"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	rest.RestConf
	Mysql   mysql.Config
	AuthRpc zrpc.RpcClientConf
}
