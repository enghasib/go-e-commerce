package middleware

import "github.com/enghasib/server/config"

type usrCtx string

const UserContextKey usrCtx = "user"

type Middlewares struct {
	cnf *config.Config
}

func NewMiddlewares(config *config.Config) *Middlewares {
	return &Middlewares{
		cnf: config,
	}
}
