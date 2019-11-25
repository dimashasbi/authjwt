package adapter

import (
	"AuthorizationJWT/engine"
	"AuthorizationJWT/model"
	"encoding/json"
	"net/http"
)

type (
	token struct {
		engine.Token
	}
)

// GetToken use for Get Token
func (t *token) CreateToken(w http.ResponseWriter, r *http.Request) {
	mod := model.Users{}
	json.NewDecoder(r.Body).Decode(&mod)

	result := t.CreateTokenUsecase(mod)

	resp, _ := json.Marshal(result)

	DefaultRespon(w, resp)
}

// CheckToken use for Check Token
func (t *token) CheckToken(w http.ResponseWriter, r *http.Request) {
	mod := model.TokenCookiesJwt{}
	json.NewDecoder(r.Body).Decode(&mod)

	result := t.CheckTokenUsecase(mod)

	resp, _ := json.Marshal(result)

	DefaultRespon(w, resp)
}
