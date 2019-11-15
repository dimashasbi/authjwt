package adapter

import (
	"net/http"
)

// ErrRespon struct if get Error Application
type ErrRespon struct {
	Error string
}

// DefaultRespon use for no Respon Build
func DefaultRespon(w http.ResponseWriter, resp []byte) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(resp))
}
