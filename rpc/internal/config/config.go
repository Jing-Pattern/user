package config

import "github.com/zeromicro/go-zero/zrpc"

type Mysql struct {
	DbUrl string
}
type Config struct {
	zrpc.RpcServerConf
	Mysql Mysql
}
