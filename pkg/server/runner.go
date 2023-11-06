package server

import "go.uber.org/zap"

// Run runs the given application.
func Run(app Application) error {
	app.logger.Debug("Listening...", zap.String("address", app.listener.Addr().String()))
	if err := app.Run(); err != nil {
		return err
	}
	return app.Shutdown()
}
