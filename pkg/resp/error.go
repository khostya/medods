package resp

import (
	"github.com/go-chi/render"
	"medods/pkg/validator"
	"net/http"
)

type RespError struct {
	Type    string `json:"type"`
	Message string `json:"message"`

	ValidationError *validator.ValidationError `json:"validationError"`
}

func Error(r *http.Request, w http.ResponseWriter, err error, status int) {
	if err == nil {
		return
	}

	render.Status(r, status)
	switch err.(type) {
	case validator.ValidationError:
		validationError(r, w, err.(validator.ValidationError))
	default:
		rawError(r, w, err)
	}
}

func validationError(r *http.Request, w http.ResponseWriter, err validator.ValidationError) {
	resp := RespError{Type: "validation", ValidationError: &err}
	render.JSON(w, r, resp)
}

func rawError(r *http.Request, w http.ResponseWriter, err error) {
	resp := RespError{Type: "raw", Message: err.Error()}
	render.JSON(w, r, resp)
}
