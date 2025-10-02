package product

import (
	"net/http"

	middleware "github.com/enghasib/server/rest/middlewares"
)

func (h *ProductHandler) ProductRoutes(mux *http.ServeMux, manager *middleware.Manager) *http.ServeMux {
	// products routes
	mux.Handle("GET /products", manager.With(http.HandlerFunc(h.GetProductHandler)))
	mux.Handle("GET /products/{id}", manager.With(http.HandlerFunc(h.GetSingleProduct)))

	mux.Handle("POST /products",
		manager.With(http.HandlerFunc(
			h.CreateProductHandler,
		), h.middlewares.Authentication),
	)

	mux.Handle("PATCH /products/{id}",
		manager.With(http.HandlerFunc(
			h.UpdateProductHandler,
		), h.middlewares.Authentication),
	)

	mux.Handle("DELETE /products/{id}", manager.With(http.HandlerFunc(h.DeleteProductHandler)))

	return mux
}
