package app

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/drizzleent/patients/internal/config"
)

type App struct {
	serviceProvider *serviceProvider
	httpServer      *http.Server
}

func New(ctx context.Context) (*App, error) {
	a := &App{}
	err := a.initDebs(ctx)
	if err != nil {
		return nil, err
	}
	return a, nil
}

func (a *App) Run() error {
	err := a.runHttpServer()
	if err != nil {
		return err
	}

	return nil
}

func (a *App) initDebs(ctx context.Context) error {
	inits := []func(ctx context.Context) error{
		a.initConfig,
		a.initServiceProvider,
		a.initHttpServer,
		a.initFile,
	}

	for _, f := range inits {
		err := f(ctx)
		if err != nil {
			return err
		}
	}
	return nil
}

func (a *App) initServiceProvider(_ context.Context) error {
	a.serviceProvider = newServiceProvider()
	return nil
}

func (a *App) initConfig(_ context.Context) error {
	err := config.Load(".env")
	if err != nil {
		return err
	}

	return nil
}
func (a *App) initFile(ctx context.Context) error {
	return a.serviceProvider.Repository(ctx).Load()
}

func (a *App) initHttpServer(ctx context.Context) error {
	srv := &http.Server{
		Addr:           a.serviceProvider.HTTPConfig().Address(),
		Handler:        a.serviceProvider.Handler(ctx).InitRoutes(),
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	a.httpServer = srv
	return nil
}

func (a *App) runHttpServer() error {
	log.Printf("server run on %s", a.serviceProvider.HTTPConfig().Address())
	err := a.httpServer.ListenAndServe()

	if err != nil {
		return err
	}

	return nil
}
