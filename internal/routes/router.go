package routes

import (
	"github.com/joydebnath/bebolt/internal"
	"github.com/uptrace/bunrouter"
	"github.com/uptrace/bunrouter/extra/reqlog"
)

type Router struct {
	Router *bunrouter.Router
	App    *internal.App
}

func NewRouter(app *internal.App) *Router {
	return &Router{
		App: app,
		Router: bunrouter.New(
			bunrouter.Use(reqlog.NewMiddleware(
				reqlog.WithEnabled(false),
				reqlog.FromEnv("BUNDEBUG"),
			)),
		),
	}
}
