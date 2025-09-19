package user

import (
	"encoding/json"
	"net/http"

	"github.com/enghasib/server/internal/models"
)

type UserResponse struct {
	Message     string `json:"message"`
	ID          int    `json:"id"`
	UserName    string `json:"user_name"`
	Email       string `json:"email"`
	IsShopOwner bool   `json:"is_shop_owner"`
}

// create product
func (h *UserHandler) CreateUserHandler(w http.ResponseWriter, r *http.Request) {

	var newUser models.User

	err := json.NewDecoder(r.Body).Decode(&newUser)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	newUser.ID = len(models.UserList) + 1
	models.UserList = append(models.UserList, newUser)

	// mount and encode with response
	json.NewEncoder(w).Encode(UserResponse{
		Message:     "User created Successfully!",
		ID:          newUser.ID,
		UserName:    newUser.UserName,
		Email:       newUser.Email,
		IsShopOwner: newUser.IsShowOwner,
	})

}
