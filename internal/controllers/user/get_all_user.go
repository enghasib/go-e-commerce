package user

import (
	"encoding/json"
	"net/http"

	"github.com/enghasib/server/internal/models"
)

func (h *UserHandler) GetAllUserHandler(w http.ResponseWriter, r *http.Request) {
	// if r.Method != http.MethodGet {
	// 	http.Error(w, "Only accept GET request!", http.StatusBadRequest)
	// 	return
	// }
	err := json.NewEncoder(w).Encode(models.UserList)
	if err != nil {
		panic(err)
	}
}
