package app

import (
	"context"
	"log"

	"github.com/drizzleent/patients/internal/api"
	"github.com/drizzleent/patients/internal/api/http/handler"
	"github.com/drizzleent/patients/internal/config"
	"github.com/drizzleent/patients/internal/config/env"
	"github.com/drizzleent/patients/internal/repository"
	patientsRepo "github.com/drizzleent/patients/internal/repository/patients"
	"github.com/drizzleent/patients/internal/service"
	patientsService "github.com/drizzleent/patients/internal/service/patients"
)

type serviceProvider struct {
	httpConfig config.HTTPConfig

	repository repository.Repository

	service service.ApiService

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

func (s *serviceProvider) Repository(ctx context.Context) repository.Repository {
	if nil == s.repository {
		s.repository = patientsRepo.NewRepository()
	}

	return s.repository
}

func (s *serviceProvider) Service(ctx context.Context) service.ApiService {
	if nil == s.service {
		s.service = patientsService.NewService(s.Repository(ctx))
	}

	return s.service
}

func (s *serviceProvider) Handler(ctx context.Context) api.Handler {
	if nil == s.handler {
		s.handler = handler.NewHandler(s.Service(ctx))
	}

	return s.handler
}
