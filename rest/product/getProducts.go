package product

import (
	"net/http"
	"strconv"

	"github.com/enghasib/server/domain"
	"github.com/enghasib/server/utils"
)

type pagination struct {
	Data []*domain.Product `json:"data"`
}

func (h *ProductHandler) GetProductHandler(w http.ResponseWriter, r *http.Request) {

	query := r.URL.Query()
	limit, _ := strconv.Atoi(query.Get("limit"))
	if limit == 0 || limit > 100 {
		limit = 10
	}
	page, _ := strconv.Atoi(query.Get("page"))
	if page == 0 {
		page = 1
	}

	productList, err := h.srv.List(limit, page)
	if err != nil {
		utils.SendErrorWithError(w, http.StatusInternalServerError, "Server error!")
	}

	countData, err := h.srv.Count()
	if err != nil {
		utils.SendErrorWithError(w, http.StatusInternalServerError, "Server error!")
	}

	utils.SendResponseWithPagination(w, productList, page, limit, countData)

}
