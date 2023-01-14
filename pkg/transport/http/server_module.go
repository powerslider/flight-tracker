package http

import (
	"context"
	"github.com/pkg/errors"
	"go.uber.org/fx"
)

func ServerModule() fx.Option {
	return fx.Options(
		fx.Provide(
			NewServer,
		),
	)
}

func RunServer(lc fx.Lifecycle, server *Server, errChan chan error) {
	lc.Append(fx.Hook{
		OnStart: onStart(errChan, server),
		OnStop:  onStop(server),
	})
}

func onStart(errChan chan error, srv *Server) func(context.Context) error {
	return func(ctx context.Context) error {
		go srv.Start(ctx, errChan)

		return nil
	}
}

func onStop(srv *Server) func(context.Context) error {
	return func(ctx context.Context) error {
		err := srv.Stop(ctx)
		if err != nil {
			return errors.WithStack(err)
		}

		return nil
	}
}
