package cmd

import (
	"fmt"
	"net/http"

	middleware "github.com/enghasib/server/internal/middlewares"
	"github.com/enghasib/server/internal/routes"
)

func Serve() {

	manager := middleware.MManager()
	manager.Use(middleware.Logger, middleware.Test)

	mux := http.NewServeMux()
	routes := routes.Routes(mux)

	fmt.Println("Server is running on port 4000")
	err := http.ListenAndServe(":4000", manager.Apply(routes))
	if err != nil {
		panic(err)
	}
}
