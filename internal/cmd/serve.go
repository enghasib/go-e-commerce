package cmd

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/enghasib/server/internal/config"
	"github.com/enghasib/server/internal/controllers/product"
	"github.com/enghasib/server/internal/controllers/user"
	middleware "github.com/enghasib/server/internal/middlewares"
)

type server struct {
	UserHandler    *user.UserHandler
	ProductHandler *product.ProductHandler
}

func NewServer(
	userHandler *user.UserHandler,
	productHandler *product.ProductHandler,
) *server {
	return &server{
		UserHandler:    userHandler,
		ProductHandler: productHandler,
	}
}

func (s *server) Serve() {
	config.LoadEnv()
	mux := http.NewServeMux()

	manager := middleware.MManager()
	manager.Use(middleware.Logger, middleware.Test, middleware.Cors)

	s.UserHandler.UserRoute(mux)
	s.ProductHandler.ProductRoutes(mux)

	wrappedMux := manager.Apply(mux)

	fmt.Println("Server is running on port:", config.Env.HttpPort)

	port := ":" + strconv.FormatUint(uint64(config.Env.HttpPort), 10)
	err := http.ListenAndServe(port, wrappedMux)
	if err != nil {
		panic(err)
	}
}
