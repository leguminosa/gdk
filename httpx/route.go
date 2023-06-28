package httpx

import "net/http"

// GET calls mux HandleFunc method with GET method and given handler func
func (h *Router) GET(path string, f func(http.ResponseWriter, *http.Request)) {
	h.handleFunc(methodGet, path, f)
}

// POST calls mux HandleFunc method with POST method and given handler func
func (h *Router) POST(path string, f func(http.ResponseWriter, *http.Request)) {
	h.handleFunc(methodPost, path, f)
}

// PUT calls mux HandleFunc method with PUT method and given handler func
func (h *Router) PUT(path string, f func(http.ResponseWriter, *http.Request)) {
	h.handleFunc(methodPut, path, f)
}

// DELETE calls mux HandleFunc method with DELETE method and given handler func
func (h *Router) DELETE(path string, f func(http.ResponseWriter, *http.Request)) {
	h.handleFunc(methodDelete, path, f)
}

// HandleFunc calls mux HandleFunc method with given method and handler func
func (h *Router) HandleFunc(method, path string, f func(http.ResponseWriter, *http.Request)) {
	h.handleFunc(method, path, f)
}

func (h *Router) handleFunc(method, path string, f func(http.ResponseWriter, *http.Request)) {
	_ = h.mux.HandleFunc(path, newHandler(method, f).ServeHTTP)
}
