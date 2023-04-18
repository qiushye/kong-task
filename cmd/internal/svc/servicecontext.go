package svc

import (
	"kong-task/cmd/internal/config"
	"kong-task/cmd/internal/middleware"
	"kong-task/cmd/internal/model"

	"github.com/zeromicro/go-zero/rest"
)

type ServiceContext struct {
	Config    config.Config
	Auth      rest.Middleware
	RuleModel model.RuleModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:    c,
		Auth:      middleware.NewAuthMiddleware().Handle,
		RuleModel: model.NewRuleModel(c.Mongo.Url, c.Mongo.Url, c.Redis),
	}
}
