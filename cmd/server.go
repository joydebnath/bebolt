package cmd

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"time"

	"github.com/joydebnath/bebolt/internal"
	"github.com/joydebnath/bebolt/internal/routes"
	"github.com/klauspost/compress/gzhttp"
	"github.com/rs/cors"
	"go.uber.org/fx"
)

func NewServer(lc fx.Lifecycle) *http.Server {
	app := internal.NewApp()
	r := routes.NewRouter(app)
	r.SetWebRoutes()
	r.SetAPIRoutes()

	handler := http.Handler(r.Router)
	handler = cors.Default().Handler(handler)
	handler = gzhttp.GzipHandler(handler)

	srvr := http.Server{
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  60 * time.Second,
		Handler:      handler,
	}

	addr := fmt.Sprintf("%s:%s", app.Env.Host, app.Env.Port)
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			ln, err := net.Listen("tcp", addr)
			if err != nil {
				return err
			}
			app.Logger.Infof("listening on %s\n", addr)
			go func() {
				err := srvr.Serve(ln)
				if err != nil {
					app.Logger.Error(err)
				}
			}()
			app.Logger.Infoln("Press CTRL+C to exit...")
			return nil
		},
		OnStop: func(ctx context.Context) error {
			app.Logger.Warnln("Shutting down server...")
			if err := srvr.Shutdown(ctx); err != nil {
				app.Logger.Error(err)
			}
			app.Logger.Warnln("Server is DOWN")
			return nil
		},
	})

	return &srvr
}
