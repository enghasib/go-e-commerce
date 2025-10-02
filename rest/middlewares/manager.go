package middleware

import (
	"net/http"
)

type Middleware func(http.Handler) http.Handler

type Manager struct {
	globalMiddlewares []Middleware
}

func NewManager() *Manager {
	return &Manager{
		globalMiddlewares: make([]Middleware, 0),
	}
}

func (mng *Manager) Use(middleware ...Middleware) {
	mng.globalMiddlewares = append(mng.globalMiddlewares, middleware...)

}

func (mng *Manager) Apply(handler http.Handler) http.Handler {
	for i := len(mng.globalMiddlewares) - 1; i >= 0; i-- {
		handler = mng.globalMiddlewares[i](handler)
	}
	return handler
}

func (mng *Manager) With(handler http.Handler, middlewares ...Middleware) http.Handler {

	for i := len(middlewares) - 1; i >= 0; i-- {
		handler = middlewares[i](handler)
	}
	return handler
}
