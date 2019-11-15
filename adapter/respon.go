package adapter

import (
	"net/http"
)

// DefaultRespon use for no Respon Build
func DefaultRespon(w http.ResponseWriter, resp []byte) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(resp))
}
