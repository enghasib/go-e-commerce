package user

import (
	"github.com/enghasib/server/config"
	middleware "github.com/enghasib/server/rest/middlewares"
)

type UserHandler struct {
	middleware middleware.Middlewares
	cnf        *config.Config
	srv        Service
}

func NewUserHandler(
	middlewares *middleware.Middlewares,
	cnf *config.Config,
	srv Service,
) *UserHandler {
	return &UserHandler{
		middleware: *middlewares,
		cnf:        cnf,
		srv:        srv,
	}
}
