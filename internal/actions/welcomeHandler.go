package actions

import (
	"net/http"

	"github.com/uptrace/bunrouter"
)

func (a *Action) WelcomeHandler(w http.ResponseWriter, req bunrouter.Request) error {
	a.app.Logger.Infoln("Logging from welcome handler")
	return bunrouter.JSON(w, bunrouter.H{
		"message": "Hello from Go API!",
	})
}
