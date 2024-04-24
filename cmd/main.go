package main

import (
	"context"
	"log"

	"github.com/drizzleent/patients/internal/app"
)

func main() {
	ctx := context.Background()
	app, err := app.New(ctx)
	if err != nil {
		log.Fatalf("failed to init app: %v", err.Error())
	}

	err = app.Run()
	if err != nil {
		log.Fatalf("failed to run app: %v", err.Error())
	}
}
