package user

import (
	"encoding/json"
	"net/http"

	"github.com/enghasib/server/internal/repo"
)

type UserResponse struct {
	Message     string `json:"message"`
	ID          int    `json:"id"`
	UserName    string `json:"user_name"`
	Email       string `json:"email"`
	IsShopOwner bool   `json:"is_shop_owner"`
}

type createUserRequest struct {
	UserName    string `json:"user_name"`
	Email       string `json:"email"`
	Password    string `json:"password"`
	IsShowOwner bool   `json:"is_shop_owner"`
}

// create product
func (h *UserHandler) CreateUserHandler(w http.ResponseWriter, r *http.Request) {

	var newUser createUserRequest

	err := json.NewDecoder(r.Body).Decode(&newUser)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	user, err := h.userRepo.Create(repo.User{
		UserName:    newUser.UserName,
		Email:       newUser.Email,
		Password:    newUser.Password,
		IsShowOwner: newUser.IsShowOwner,
	})

	if err != nil {
		http.Error(w, "user creation failed", http.StatusBadRequest)
		return
	}

	// mount and encode with response
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(UserResponse{
		Message:     "User created Successfully!",
		ID:          user.ID,
		UserName:    user.UserName,
		Email:       user.Email,
		IsShopOwner: user.IsShowOwner,
	})

}
