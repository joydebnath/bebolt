package main

import (
	"net/http"

	"github.com/joydebnath/bebolt/cmd"
	"go.uber.org/fx"
)

func main() {
	fx.New(
		fx.Provide(cmd.NewServer),
		fx.Invoke(func(*http.Server) {}),
	).Run()
}
