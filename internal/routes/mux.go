package routes

import (
	"fmt"
	"net/http"

	productController "github.com/enghasib/server/internal/controllers/product"
	middleware "github.com/enghasib/server/internal/middlewares"
)

func Routes(mux *http.ServeMux) *http.ServeMux {

	mux.Handle("GET /products", middleware.With(http.HandlerFunc(productController.GetProductHandler)))
	mux.Handle("GET /products/{id}", middleware.With(http.HandlerFunc(productController.GetSingleProduct)))

	mux.Handle("POST /products", http.HandlerFunc(productController.CreateProductHandler))
	mux.Handle("PATCH /products/{id}", http.HandlerFunc(productController.UpdateProductHandler))
	mux.Handle("DELETE /products/{id}", http.HandlerFunc(productController.DeleteProductHandler))

	mux.Handle("GET /test", middleware.With(
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			fmt.Println("Hello world")
		}), middleware.LocalMiddleware,
	))

	return mux
}
