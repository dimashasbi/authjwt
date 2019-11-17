package adapter

import (
	"AuthorizationJWT/engine"
	"net/http"
)

type (
	token struct {
		engine.Token
	}
)

// GetToken use for Get Token
func (t *token) GetToken(w http.ResponseWriter, r *http.Request) {
}

// CheckToken use for Check Token
func (t *token) CheckToken(w http.ResponseWriter, r *http.Request) {
}
