package svc

import (
	"greet-custom/internal/config"
)

type ServiceContext struct {
	Config config.Config
  Transformer transformer.Transformer     
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		Transformer: transformer.NewTransformer(rpcx.MustNewClient(c.Transform)),
	}
}
