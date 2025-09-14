package productController

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/enghasib/server/internal/models"
)

func DeleteProductHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}

	urlId := r.PathValue("id")
	id, err := strconv.Atoi(urlId)

	if err != nil {
		fmt.Println("Invalid URL")
		return
	}

	index := -1

	for i, product := range models.ProductList {
		if product.ID == id {
			index = i
			break
		}
	}

	if index == -1 {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "Product not found",
		})
		return
	}

	models.ProductList = append(models.ProductList[:index], models.ProductList[index+1:]...)
	json.NewEncoder(w).Encode(&models.ProductList)

}
