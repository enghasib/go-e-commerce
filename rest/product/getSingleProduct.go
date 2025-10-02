package product

import (
	"encoding/json"
	"net/http"
	"strconv"
)

func (h *ProductHandler) GetSingleProduct(w http.ResponseWriter, r *http.Request) {

	idPram, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		http.Error(w, "invalid parameter", http.StatusBadRequest)
	}

	product, err := h.srv.Get(idPram)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
	}

	json.NewEncoder(w).Encode(product)

}
