package user

import (
	"net/http"
	"strconv"

	"github.com/enghasib/server/utils"
)

func (h *UserHandler) GetAllUserHandler(w http.ResponseWriter, r *http.Request) {

	queryParam := r.URL.Query()
	page, _ := strconv.Atoi(queryParam.Get("page"))
	if page <= 0 {
		page = 1
	}
	limit, _ := strconv.Atoi(queryParam.Get("limit"))
	if limit <= 0 || limit > 100 {
		limit = 10
	}

	userList, err := h.srv.List(limit, page)
	if err != nil {
		utils.SendErrorWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.SendResponseWithPagination(w, userList, page, limit, 0)
}
