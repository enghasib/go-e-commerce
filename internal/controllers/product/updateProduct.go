package productController

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/enghasib/server/internal/models"
)

// update struct
type ProductUpdate struct {
	Title       *string  `json:"title,omitempty"`
	Description *string  `json:"description,omitempty"`
	Price       *float64 `json:"price,omitempty"`
	ImgUrl      *string  `json:"img_url,omitempty"`
}

func UpdateProductHandler(w http.ResponseWriter, r *http.Request) {
	HandleHeaders(w)
	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}

	var updatedProduct ProductUpdate

	err := json.NewDecoder(r.Body).Decode(&updatedProduct)
	if err != nil {
		fmt.Println(err)
	}

	strId := r.PathValue("id")
	id, _ := strconv.Atoi(strId)

	var productList = models.ProductList

	for i := range productList {
		if productList[i].ID == id {
			if updatedProduct.Title != nil {
				productList[i].Title = *updatedProduct.Title
			}
			if updatedProduct.Description != nil {
				productList[i].Description = *updatedProduct.Description
			}
			if updatedProduct.Price != nil {
				productList[i].Price = *updatedProduct.Price
			}
			if updatedProduct.ImgUrl != nil {
				productList[i].ImgUrl = *updatedProduct.ImgUrl
			}
			json.NewEncoder(w).Encode(productList[i])
			return
		}
	}

}
