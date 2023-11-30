package actions

import (
	"net/http"

	"github.com/joydebnath/bebolt/internal"
	"github.com/uptrace/bunrouter"
)

type Action struct {
	app *internal.App
}

type ActionHandler interface {
	Handle(http.ResponseWriter, bunrouter.Request) error
}

func NewAction(app *internal.App) *Action {
	return &Action{
		app: app,
	}
}
