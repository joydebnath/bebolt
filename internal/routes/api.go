package routes

import (
	"net/http"

	"github.com/uptrace/bunrouter"
)

func (r *Router) SetAPIRoutes() {
	r.Router.WithGroup("/api", func(g *bunrouter.Group) {
		//api routes go here
		g.GET("/", func(w http.ResponseWriter, req bunrouter.Request) error {
			return bunrouter.JSON(w, bunrouter.H{
				"message": "api route",
			})
		})
	})
}
