package product

import (
	middleware "github.com/enghasib/server/rest/middlewares"
)

type ProductHandler struct {
	middlewares middleware.Middlewares
	srv         Service
}

func NewProductHandler(middlewares *middleware.Middlewares, srv Service) *ProductHandler {
	return &ProductHandler{
		middlewares: *middlewares,
		srv:         srv,
	}
}
