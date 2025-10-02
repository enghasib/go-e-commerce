package product

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

func (h *ProductHandler) DeleteProductHandler(w http.ResponseWriter, r *http.Request) {

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

	if err := h.srv.Delete(id); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	json.NewEncoder(w).Encode(`product delete successfully!`)
}
