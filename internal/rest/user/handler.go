package user

import (
	"github.com/enghasib/server/internal/config"
	"github.com/enghasib/server/internal/repo"
	middleware "github.com/enghasib/server/internal/rest/middlewares"
)

type UserHandler struct {
	middleware middleware.Middlewares
	cnf        *config.Config
	userRepo   repo.UserRepo
}

func NewUserHandler(
	middlewares *middleware.Middlewares,
	cnf *config.Config,
	userRepo repo.UserRepo,
) *UserHandler {
	return &UserHandler{
		middleware: *middlewares,
		cnf:        cnf,
		userRepo:   userRepo,
	}
}
