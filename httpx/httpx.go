package httpx

import (
	"net/http"

	"github.com/gorilla/mux"
)

type (
	// Router encapsulates and extends mux router functionality
	Router struct {
		mux *mux.Router
	}
)

// NewRouter creates a new instance with mux router wrapped inside
func NewRouter() *Router {
	return &Router{
		mux: mux.NewRouter(),
	}
}

// ServeHTTP calls mux implementation of ServeHTTP
func (h *Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.mux.ServeHTTP(w, r)
}
