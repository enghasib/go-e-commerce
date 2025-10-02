package user

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

func (h *UserHandler) GetAllUserHandler(w http.ResponseWriter, r *http.Request) {
	// if r.Method != http.MethodGet {
	// 	http.Error(w, "Only accept GET request!", http.StatusBadRequest)
	// 	return
	// }

	userList, err := h.srv.List()
	if err != nil {
		http.Error(w, "no user found", http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(userList)
	if err != nil {
		fmt.Println("error:", err)
		os.Exit(1)
	}
}
