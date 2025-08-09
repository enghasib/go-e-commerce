package cmd

import (
	"fmt"
	"net/http"

	middleware "github.com/enghasib/server/internal/middlewares"
	mux "github.com/enghasib/server/internal/routes"
)

func Serve() {
	mux := mux.MainMux()

	fmt.Println("Server is running on port 4000")
	err := http.ListenAndServe(":4000", middleware.Logger(mux))
	if err != nil {
		panic(err)
	}
}
