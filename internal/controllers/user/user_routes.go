package user

import (
	"net/http"

	middleware "github.com/enghasib/server/internal/middlewares"
)

func (h *UserHandler) UserRoute(mux *http.ServeMux) *http.ServeMux {
	mux.Handle("POST /users", http.HandlerFunc(h.CreateUserHandler))
	mux.Handle("POST /users/login", http.HandlerFunc(h.LoginUser))

	mux.Handle("GET /users", middleware.With(http.HandlerFunc(h.GetAllUserHandler)))

	return mux
}
