package user

import (
	"encoding/json"
	"net/http"

	"github.com/enghasib/server/domain"
	"github.com/enghasib/server/utils"
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
		utils.SendErrorWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	user, err := h.srv.Create(domain.User{
		UserName:    newUser.UserName,
		Email:       newUser.Email,
		Password:    newUser.Password,
		IsShowOwner: newUser.IsShowOwner,
	})

	if err != nil {
		utils.SendErrorWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	// mount and encode with response
	utils.SendResponseWithData(w, http.StatusCreated, UserResponse{
		Message:     "User created Successfully!",
		ID:          user.ID,
		UserName:    user.UserName,
		Email:       user.Email,
		IsShopOwner: user.IsShowOwner,
	})

}
