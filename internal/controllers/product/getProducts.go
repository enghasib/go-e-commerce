package productController

import (
	"encoding/json"
	"net/http"

	"github.com/enghasib/server/internal/models"
)

func GetProductHandler(w http.ResponseWriter, r *http.Request) {
	HandleHeaders(w)
	if r.Method != http.MethodGet {
		http.Error(w, "Only accept GET request!", http.StatusBadRequest)
		return
	}
	err := json.NewEncoder(w).Encode(models.ProductList)
	if err != nil {
		panic(err)
	}
}
