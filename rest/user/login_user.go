package user

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/enghasib/server/utils"
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

	user, err := h.srv.Find(requestBody.Email, requestBody.Password)
	if err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	payload := utils.Payload{
		Sub:         user.ID,
		UserName:    user.UserName,
		Email:       user.Email,
		IsShopOwner: user.IsShowOwner,
	}

	jwt, err := utils.CreateToken(h.cnf.JwtSecret, payload)
	if err != nil {
		utils.SendErrorWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	response := loginResponse{
		Message:     "User Login successfully!",
		AccessToken: jwt,
	}

	utils.SendResponseWithData(w, http.StatusOK, response)
}
