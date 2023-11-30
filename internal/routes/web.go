package routes

import (
	"github.com/joydebnath/bebolt/internal/actions"
)

func (r *Router) SetWebRoutes() {
	//web routes go here
	a := actions.NewAction(r.App)
	r.Router.GET("/", a.WelcomeHandler)
}
