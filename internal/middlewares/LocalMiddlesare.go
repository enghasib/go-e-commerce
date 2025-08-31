package middleware

import (
	"fmt"
	"net/http"
)

func LocalMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Request header:", r.Header)
		fmt.Println("Request Length:", r.ContentLength)
		next.ServeHTTP(w, r)
	})
}
