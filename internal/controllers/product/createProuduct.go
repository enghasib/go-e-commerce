package product

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/enghasib/server/internal/models"
)

type Response struct {
	Message string `json:"message"`
	Product models.Product
}

// create product
func (h *ProductHandler) CreateProductHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Create product handler call .......")
	//extract body and decode
	// var productList = models.ProductList
	var newProduct models.Product
	err := json.NewDecoder(r.Body).Decode(&newProduct)
	if err != nil {
		fmt.Println("Json parsing error:", err.Error())
	}

	newProduct.ID = len(models.ProductList) + 1

	models.ProductList = append(models.ProductList, newProduct)

	response := Response{
		Message: "User created successfully!",
		Product: newProduct,
	}
	// mount and encode with response
	json.NewEncoder(w).Encode(response)

}
