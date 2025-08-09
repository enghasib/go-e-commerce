package productController

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/enghasib/server/internal/models"
)

// create product
func CreateProductHandler(w http.ResponseWriter, r *http.Request) {
	HandleHeaders(w)
	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}

	//extract body and decode
	var productList = models.ProductList
	var newProduct models.Product
	err := json.NewDecoder(r.Body).Decode(&newProduct)
	if err != nil {
		fmt.Println("Json parsing error:", err.Error())
	}

	newProduct.ID = len(models.ProductList) + 1

	models.ProductList = append(models.ProductList, newProduct)

	// mount and encode with response
	json.NewEncoder(w).Encode(productList)

}
