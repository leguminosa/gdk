package httpx

import "github.com/gorilla/mux"

// Use registers multiple mux middlewares by iterating each of them
func (h *Router) Use(mwf ...mux.MiddlewareFunc) {
	h.mux.Use(mwf...)
}
