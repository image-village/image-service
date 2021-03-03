package middlewares

import (
	"errors"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"github.com/lagbana/images/server/auth"
	"github.com/lagbana/images/server/responses"
	"net/http"
)

// Middleware - wrapper around original app handler
type Middleware struct {
	next http.Handler
}

// GlobalAuthCheck - Constructor for Application level middleware since the fields are not exported
func GlobalAuthCheck(next http.Handler) *Middleware {
	return &Middleware{next: next}
}

// ServeHttp - Application level middleware handler
func (m *Middleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Check headers for token, verify if found and return payload
	payload, err := auth.ValidateToken(r)
	if err != nil || payload == "" {
		fmt.Println(`❌`, err)
		m.next.ServeHTTP(w, r)
		return
	}

	// Set payload to Header
	r.Header.Set("currentUser", payload)
	m.next.ServeHTTP(w, r)
}

// AuthGuard - Route level middleware 
func AuthGuard(next httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		v := r.Header.Get("currentUser")
		if v == "" {
			responses.ERROR(w, http.StatusUnauthorized, errors.New("Unauthorized"))
			return
		}

		next(w, r, p)
	}
}
