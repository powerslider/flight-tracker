package main

import (
	"context"
	"log"

	"github.com/powerslider/flight-tracker/pkg/configs"
	"github.com/powerslider/flight-tracker/pkg/handlers"
	httptransport "github.com/powerslider/flight-tracker/pkg/transport/http"
	"go.uber.org/fx"
)

// @title Flight Tracker API
// @version 1.0
// @description Demo service for tracking flights.
// @termsOfService http://swagger.io/terms/

// @contact.name Tsvetan Dimitrov
// @contact.email tsvetan.dimitrov23@gmail.com

// @license.name MIT
// @license.url https://www.mit.edu/~amini/LICENSE.md

// @host localhost:8080
// @BasePath /
func main() {
	errCh := make(chan error)
	ctx, cancel := context.WithCancel(context.Background())

	defer cancel()

	app := fx.New(
		fx.Supply(errCh),
		fx.Provide(
			func() context.Context {
				return ctx
			},
		),
		configs.Module(),
		httptransport.ServerModule(),
		handlers.Module(),

		fx.Invoke(httptransport.RunServer),
	)

	if err := app.Start(ctx); err != nil {
		panic(err)
	}

	select {
	case <-ctx.Done():
		log.Println(ctx, "Context cancelled. Exiting...")
	case <-app.Done():
		log.Println(ctx, "Interrupt received. Exiting...")

		if err := app.Stop(ctx); err != nil {
			panic(err)
		}
	case err := <-errCh:
		log.Println(ctx, "App errored:", err)
	}
}
