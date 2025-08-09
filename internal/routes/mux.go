package mux

import (
	"net/http"

	productController "github.com/enghasib/server/internal/controllers/product"
)

func MainMux() *http.ServeMux {
	mux := http.NewServeMux()

	mux.Handle("GET /products", http.HandlerFunc(productController.GetProductHandler))
	mux.Handle("POST /products", http.HandlerFunc(productController.CreateProductHandler))
	mux.Handle("PATCH /products/{id}", http.HandlerFunc(productController.UpdateProductHandler))

	return mux
}
