package httpx

import "net/http"

type (
	handler struct {
		method  string
		handler http.Handler
	}
)

func newHandler(method string, f func(http.ResponseWriter, *http.Request)) *handler {
	return &handler{
		method:  method,
		handler: http.HandlerFunc(f),
	}
}

// ServeHTTP determines whether the handler is eligible to be called
func (h *handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var handler http.Handler

	handler = h.handler

	// set as 405 if request method does not match
	if r.Method != h.method {
		handler = methodNotAllowedHandler()
	}

	// set as 404 if there is no appropriate handler
	if handler == nil {
		handler = notFoundHandler()
	}

	handler.ServeHTTP(w, r)
}
