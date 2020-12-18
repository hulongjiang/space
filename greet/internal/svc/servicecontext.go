package svc

import (
	"github.com/tal-tech/go-zero/zrpc"
	"space/greet/internal/config"
	"space/user/userclient"
)

type ServiceContext struct {
	Config  config.Config
	UserRpc userclient.User
}

func NewServiceContext(c config.Config) *ServiceContext {
	ur := userclient.NewUser(zrpc.MustNewClient(c.UserRpc))
	return &ServiceContext{
		Config:  c,
		UserRpc: ur,
	}
}
