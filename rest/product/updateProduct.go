package product

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/enghasib/server/domain"
	"github.com/enghasib/server/utils"
)

// update struct
type reqProductUpdateData struct {
	Title       string  `json:"title,omitempty"`
	Description string  `json:"description,omitempty"`
	Price       float64 `json:"price,omitempty"`
	ImgUrl      string  `json:"img_url,omitempty"`
}

func (h *ProductHandler) UpdateProductHandler(w http.ResponseWriter, r *http.Request) {
	// if r.Method == http.MethodOptions {
	// 	w.WriteHeader(http.StatusOK)
	// 	return
	// }

	var updatedProduct reqProductUpdateData

	err := json.NewDecoder(r.Body).Decode(&updatedProduct)
	if err != nil {
		utils.SendErrorWithError(w, http.StatusBadRequest, "invalid request body!")
	}

	strId := r.PathValue("id")
	prodId, _ := strconv.Atoi(strId)

	product, err := h.srv.Update(prodId, domain.Product{
		Title:       updatedProduct.Title,
		Description: updatedProduct.Description,
		Price:       updatedProduct.Price,
		ImgUrl:      updatedProduct.ImgUrl,
	})

	if err != nil {
		utils.SendErrorWithError(w, http.StatusBadRequest, "product update failed!")
		return
	}

	utils.SendResponseWithData(w, http.StatusOK, product)

}
