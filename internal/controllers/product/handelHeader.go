package productController

import "net/http"

func HandleHeaders(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	// w.Header().Set("Access-Control-Allow-Method", http.MethodPost,)
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
}
