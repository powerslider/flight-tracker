package handlers

import (
	"github.com/gorilla/mux"
	"github.com/powerslider/flight-tracker/pkg/flights"
	"go.uber.org/fx"
)

// Module FX module function wiring internal dependencies.
func Module() fx.Option {
	return fx.Options(
		fx.Provide(
			mux.NewRouter,
			flights.NewTracker,
			NewTrackerHandler,
		),
		fx.Invoke(registerHTTPRoutes),
	)
}
