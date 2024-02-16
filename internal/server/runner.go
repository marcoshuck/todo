package server

import (
	"context"
	"go.uber.org/zap"
)

// Run runs the given application.
func Run(app Application) error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	app.logger.Debug("Listening...", zap.String("address", app.listener.Addr().String()))
	if err := app.Run(ctx); err != nil {
		return err
	}
	return app.Shutdown(ctx)
}
