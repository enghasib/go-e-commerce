package product

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/enghasib/server/internal/repo"
)

type Response struct {
	Message string `json:"message"`
	Product repo.Product
}

type createProductRequestData struct {
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	ImgUrl      string  `json:"img_url"`
}

// create product
func (h *ProductHandler) CreateProductHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Create product handler call .......")
	//extract body and decode
	// var productList = models.ProductList
	var newProduct createProductRequestData
	err := json.NewDecoder(r.Body).Decode(&newProduct)
	if err != nil {
		fmt.Println("Json parsing error:", err.Error())
	}

	product, err := h.productRepo.Create(repo.Product{
		Title:       newProduct.Title,
		Description: newProduct.Description,
		Price:       newProduct.Price,
		ImgUrl:      newProduct.ImgUrl,
	})

	if err != nil {
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}

	response := Response{
		Message: "Product created successfully!",
		Product: *product,
	}
	// mount and encode with response
	json.NewEncoder(w).Encode(response)

}
