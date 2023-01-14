package configs

import "go.uber.org/fx"

// Module exports constructors as fx dependencies.
func Module() fx.Option {
	return fx.Options(
		fx.Provide(
			NewServerConfig,
		),
	)
}
