package product

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/enghasib/server/utils"
)

func (h *ProductHandler) DeleteProductHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}

	urlId := r.PathValue("id")
	id, err := strconv.Atoi(urlId)

	if err != nil {
		fmt.Println("Invalid URL")
		return
	}

	if err := h.srv.Delete(id); err != nil {
		utils.SendErrorWithError(w, http.StatusInternalServerError, err.Error())
	}

	utils.SendResponseWithData(w, http.StatusOK, "product deleted successfully!")

}
