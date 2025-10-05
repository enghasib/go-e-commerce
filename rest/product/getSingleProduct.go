package product

import (
	"net/http"
	"strconv"

	"github.com/enghasib/server/utils"
)

func (h *ProductHandler) GetSingleProduct(w http.ResponseWriter, r *http.Request) {

	idPram, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		utils.SendErrorWithError(w, http.StatusBadRequest, "invalid parameter")
	}

	product, err := h.srv.Get(idPram)
	if err != nil {
		utils.SendErrorWithError(w, http.StatusNotFound, err.Error())
	}

	utils.SendResponseWithData(w, http.StatusOK, product)
}
