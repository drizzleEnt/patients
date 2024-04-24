package app

import (
	"context"
	"log"

	"github.com/drizzleent/patients/internal/api"
	"github.com/drizzleent/patients/internal/api/http/handler"
	"github.com/drizzleent/patients/internal/config"
	"github.com/drizzleent/patients/internal/config/env"
)

type serviceProvider struct {
	httpConfig config.HTTPConfig

	handler api.Handler
}

func newServiceProvider() *serviceProvider {
	return &serviceProvider{}
}

func (s *serviceProvider) HTTPConfig() config.HTTPConfig {
	if nil == s.httpConfig {
		cfg, err := env.NewHttpConfig()
		if err != nil {
			log.Fatalf("failed to get http config: %s", err.Error())
		}
		s.httpConfig = cfg
	}

	return s.httpConfig
}

func (s *serviceProvider) Handler(context.Context) api.Handler {
	if nil == s.handler {
		s.handler = handler.NewHandler()
	}

	return s.handler
}
