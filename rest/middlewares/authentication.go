package middleware

import (
	"context"
	"log"
	"net/http"
	"strings"

	"github.com/enghasib/server/utils"
)

func (m *Middlewares) Authentication(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("Authentication middlewares call.....")
		// header
		AuthenticationHeader := r.Header.Get("Authorization")
		if AuthenticationHeader == "" {
			http.Error(w, "Unauthorized:", http.StatusUnauthorized)
			return
		}

		//split header and grep the token
		headerArr := strings.Split(AuthenticationHeader, " ")
		if len(headerArr) != 2 {
			utils.SendErrorWithError(w, http.StatusUnauthorized, "Unauthorized")
			return
		}

		jwt_token := headerArr[1]

		// verify token
		isVerified, err := utils.Verify(jwt_token, m.cnf.JwtSecret)
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}

		if !isVerified {
			utils.SendErrorWithError(w, http.StatusUnauthorized, "Unauthorized")
			return
		}

		user, err := utils.DecodeToken(jwt_token, m.cnf.JwtSecret)
		if err != nil {
			utils.SendErrorWithError(w, http.StatusInternalServerError, "failed to verify")
		}

		ctx := context.WithValue(r.Context(), UserContextKey, user)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
