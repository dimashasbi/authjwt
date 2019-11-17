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
		redis  RedisRepository
		mapper Mapper
	}
)

func (f *engineFactory) NewTokenEngines() Token {

	return &token{
		redis:  f.NewRedisRepository(),
		mapper: f.NewMapper(),
	}
}

func (t *token) Initialization() {

}

func (t *token) CheckTokenUsecase(ploadSign model.TokenCookiesJwt) interface{} {
	return model.Properties{}
}

func (t *token) CreateTokenUsecase(userData model.Users) model.TokenCookiesJwt {
	return model.TokenCookiesJwt{}
}
