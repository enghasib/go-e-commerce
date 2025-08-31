package productController

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/enghasib/server/internal/models"
)

func GetSingleProduct(w http.ResponseWriter, r *http.Request) {
	HandleHeaders(w)
	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}

	idPram, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		http.Error(w, "invalid parameter", http.StatusBadRequest)
	}

	for _, product := range models.ProductList {
		if product.ID == idPram {
			json.NewEncoder(w).Encode(product)
			return
		}
	}

}
