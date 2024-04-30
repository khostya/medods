package resp

import (
	"github.com/go-chi/render"
	"net/http"
)

func Json(w http.ResponseWriter, r *http.Request, json any, status int) {
	render.Status(r, status)
	render.JSON(w, r, json)
}
