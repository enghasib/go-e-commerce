package product

import (
	"encoding/json"
	"net/http"
)

func (h *ProductHandler) GetProductHandler(w http.ResponseWriter, r *http.Request) {
	// if r.Method != http.MethodGet {
	// 	http.Error(w, "Only accept GET request!", http.StatusBadRequest)
	// 	return
	// }

	productList, err := h.srv.List()
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
	}

	err = json.NewEncoder(w).Encode(productList)
	if err != nil {
		panic(err)
	}
}
