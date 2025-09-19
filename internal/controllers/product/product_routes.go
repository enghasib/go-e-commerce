package product

import (
	"net/http"

	middleware "github.com/enghasib/server/internal/middlewares"
)

func (h *ProductHandler) ProductRoutes(mux *http.ServeMux) *http.ServeMux {
	// products routes
	mux.Handle("GET /products", middleware.With(http.HandlerFunc(h.GetProductHandler)))
	mux.Handle("GET /products/{id}", middleware.With(http.HandlerFunc(h.GetSingleProduct)))

	mux.Handle("POST /products", middleware.With(http.HandlerFunc(h.CreateProductHandler), middleware.Authorization))
	mux.Handle("PATCH /products/{id}", middleware.With(http.HandlerFunc(h.UpdateProductHandler)))
	mux.Handle("DELETE /products/{id}", middleware.With(http.HandlerFunc(h.DeleteProductHandler)))

	return mux
}
