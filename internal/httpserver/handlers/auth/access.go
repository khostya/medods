package auth

import (
	"github.com/google/uuid"
	"medods/pkg/resp"
	"net/http"
)

// @Tags auth
// @Param        userID    query     string  true  " "
// @Success      200   {object}  model.Tokens        "created"
// @Success      400   {object}  resp.RespError        "bad request"
// @Success      500   {object}  resp.RespError        "internal server error"
// @Router       /access [get]
func (router Router) access(w http.ResponseWriter, r *http.Request) {
	var (
		ctx = r.Context()
	)

	userID := r.URL.Query().Get("userID")
	id, err := uuid.Parse(userID)
	if err != nil {
		resp.Error(r, w, err, http.StatusBadRequest)
		return
	}

	tokens, err := router.useCase.Access(ctx, id.String())
	if err != nil {
		resp.Error(r, w, err, http.StatusInternalServerError)
		return
	}

	resp.Json(w, r, tokens, http.StatusOK)
}
