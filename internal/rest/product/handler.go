package product

import (
	"github.com/enghasib/server/internal/repo"
	middleware "github.com/enghasib/server/internal/rest/middlewares"
)

type ProductHandler struct {
	middlewares middleware.Middlewares
	productRepo repo.ProductRepo
}

func NewProductHandler(middlewares *middleware.Middlewares, productRepo repo.ProductRepo) *ProductHandler {
	return &ProductHandler{
		middlewares: *middlewares,
		productRepo: productRepo,
	}
}
