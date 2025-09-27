package rest

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/enghasib/server/internal/config"
	middleware "github.com/enghasib/server/internal/rest/middlewares"
	"github.com/enghasib/server/internal/rest/product"
	"github.com/enghasib/server/internal/rest/user"
)

type server struct {
	cnf            *config.Config
	UserHandler    *user.UserHandler
	ProductHandler *product.ProductHandler
}

func NewServer(
	cnf *config.Config,
	userHandler *user.UserHandler,
	productHandler *product.ProductHandler,
) *server {
	return &server{
		cnf:            cnf,
		UserHandler:    userHandler,
		ProductHandler: productHandler,
	}
}

func (s *server) Start() {
	mux := http.NewServeMux()

	manager := middleware.NewManager()
	manager.Use(middleware.Logger, middleware.Test, middleware.Cors)

	s.UserHandler.UserRoute(mux, manager)
	s.ProductHandler.ProductRoutes(mux, manager)

	wrappedMux := manager.Apply(mux)

	fmt.Println("Server is running on port:", s.cnf.HttpPort)

	port := ":" + strconv.FormatUint(uint64(s.cnf.HttpPort), 10)
	err := http.ListenAndServe(port, wrappedMux)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
}
