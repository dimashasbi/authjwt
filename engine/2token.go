package engine

import (
	"AuthorizationJWT/model"
)

type (
	// Token is the interface for interactor
	Token interface {
		Initialization()
		CheckTokenUsecase(ploadSign model.TokenCookiesJwt) interface{}
		CreateTokenUsecase(userData model.Users) model.TokenCookiesJwt
	}

	token struct {
		key KeyRepository
	}
)

func (f *engineFactory) NewTokenEngines() Token {
	return &token{
		key: f.NewRedisRepository(),
	}
}

func (t *token) Initialization() {

}
