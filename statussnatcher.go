package statussnatcher

import "net/http"

// StatusSnatcher embeds the ResponseWriter interface and adds a StatusCode property which we can access.
type StatusSnatcher struct {
	http.ResponseWriter
	StatusCode int
}

// Override ResponseWriter's WriteHeader method.
func (ss *StatusSnatcher) WriteHeader(code int) {
	ss.ResponseWriter.WriteHeader(code)
	ss.StatusCode = code
}

// Constructor to create a new instance of StatusSnatcher wrapper around the original http.ResponseWriter.
func New(w http.ResponseWriter) *StatusSnatcher {
	return &StatusSnatcher{ResponseWriter: w}
}
