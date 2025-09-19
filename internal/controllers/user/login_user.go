package user

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/enghasib/server/internal/config"
	"github.com/enghasib/server/internal/models"
	"github.com/enghasib/server/internal/utils"
)

type LoginCredential struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type loginResponse struct {
	Message     string `json:"message"`
	AccessToken string `json:"access_token"`
}

func (h *UserHandler) LoginUser(w http.ResponseWriter, r *http.Request) {

	var requestBody LoginCredential
	err := json.NewDecoder(r.Body).Decode(&requestBody)

	if err != nil {
		fmt.Println("error:", err.Error())
		http.Error(w, "Invalid credentials", http.StatusBadRequest)
		return
	}

	var user *models.User
	for _, u := range models.UserList {
		if u.Email == requestBody.Email && u.Password == requestBody.Password {
			user = &u
			break
		}
	}

	if user == nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	payload := utils.Payload{
		Sub:         user.ID,
		UserName:    user.UserName,
		Email:       user.Email,
		IsShopOwner: user.IsShowOwner,
	}

	jwt, err := utils.CreateToken(config.Env.JwtSecret, payload)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := loginResponse{
		Message:     "User Login successfully!",
		AccessToken: jwt,
	}

	json.NewEncoder(w).Encode(response)
}
