package auth

import (
	"encoding/base64"
	"medods/pkg/decode"
	"medods/pkg/resp"
	"net/http"
)

type Refresh struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}

// @Tags auth
// @Param       request body Refresh true " "
// @Success      201   {object}  model.Tokens        "refreshed"
// @Success      400   {object}  resp.RespError        "bad request"
// @Success      500   {object}  resp.RespError        "internal server error"
// @Router       /refresh [post]
func (router Router) refresh(w http.ResponseWriter, r *http.Request) {
	var (
		refresh Refresh
		ctx     = r.Context()
	)

	err := decode.Json(r.Body, &refresh)
	if err != nil {
		resp.Error(r, w, err, http.StatusBadRequest)
		return
	}

	refreshToken, err := base64.StdEncoding.DecodeString(refresh.RefreshToken)
	if err != nil {
		resp.Error(r, w, err, http.StatusBadRequest)
		return
	}

	userID, err := router.tokenManager.ExtractUserIdWithoutClaimsValidation(refresh.AccessToken)
	if err != nil {
		resp.Error(r, w, err, http.StatusBadRequest)
		return
	}

	sessionID, err := router.tokenManager.ExtractIDWithoutClaimsValidation(refresh.AccessToken)
	if err != nil {
		resp.Error(r, w, err, http.StatusBadRequest)
		return
	}

	tokens, err := router.useCase.Refresh(ctx, sessionID, string(refreshToken), userID)
	if err != nil {
		resp.Error(r, w, err, http.StatusInternalServerError)
		return
	}

	resp.Json(w, r, tokens, http.StatusCreated)
}
