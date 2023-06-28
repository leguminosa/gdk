package httpx

import (
	"fmt"
	"net/http"
)

const (
	// http methods
	methodGet    = http.MethodGet
	methodPost   = http.MethodPost
	methodPut    = http.MethodPut
	methodDelete = http.MethodDelete

	// http statuses
	statusNotFound         = http.StatusNotFound
	statusMethodNotAllowed = http.StatusMethodNotAllowed
)

// default handlers

// 404 handler
func notFound(w http.ResponseWriter, r *http.Request) {
	writeError(w, "404 not found", statusNotFound)
}
func notFoundHandler() http.Handler { return http.HandlerFunc(notFound) }

// 405 handler
func methodNotAllowed(w http.ResponseWriter, r *http.Request) {
	writeError(w, "405 method not allowed", statusMethodNotAllowed)
}
func methodNotAllowedHandler() http.Handler { return http.HandlerFunc(methodNotAllowed) }

func writeError(w http.ResponseWriter, err string, code int) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.WriteHeader(code)
	fmt.Fprintln(w, err)
}
