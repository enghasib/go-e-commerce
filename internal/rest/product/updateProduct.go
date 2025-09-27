package product

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/enghasib/server/internal/repo"
)

// update struct
type reqProductUpdateData struct {
	Title       string  `json:"title,omitempty"`
	Description string  `json:"description,omitempty"`
	Price       float64 `json:"price,omitempty"`
	ImgUrl      string  `json:"img_url,omitempty"`
}

func (h *ProductHandler) UpdateProductHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}

	var updatedProduct reqProductUpdateData

	err := json.NewDecoder(r.Body).Decode(&updatedProduct)
	if err != nil {
		fmt.Println(err)
	}

	strId := r.PathValue("id")
	prodId, _ := strconv.Atoi(strId)

	product, err := h.productRepo.Update(prodId, repo.Product{
		Title:       updatedProduct.Title,
		Description: updatedProduct.Description,
		Price:       updatedProduct.Price,
		ImgUrl:      updatedProduct.ImgUrl,
	})

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(product)

}
