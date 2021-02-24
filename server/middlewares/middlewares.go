package middlewares

import (
	// "errors"
	"net/http"
	"github.com/julienschmidt/httprouter"
)

// SetMiddlewareJSON -
func SetMiddlewareJSON(next httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		w.Header().Set("Content-Type", "application/json")
		next(w, r, p)
	}
}