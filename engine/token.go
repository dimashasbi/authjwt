package engine

import (
	"AuthorizationJWT/model"
	"sync"
)

type (
	// Token is the interface for interactor
	Token interface {
		Initialization(wg *sync.WaitGroup)
		CheckTokenUsecase(ploadSign model.TokenCookiesJwt, wg *sync.WaitGroup) interface{}
		CreateTokenUsecase(userData model.Users, wg *sync.WaitGroup) model.TokenCookiesJwt
	}

	token struct {
		redis  RedisRepository
		mapper Mapper
	}
)

func (f *engineFactory) NewToken() Token {

	return &token{
		redis:  f.NewRedisRepository(),
		mapper: f.NewMapper(),
	}
}

func (t *token) Initialization(wg *sync.WaitGroup) {

}

func (t *token) CheckTokenUsecase(ploadSign model.TokenCookiesJwt, wg *sync.WaitGroup) interface{} {
	return model.Properties{}
}

func (t *token) CreateTokenUsecase(userData model.Users, wg *sync.WaitGroup) model.TokenCookiesJwt {
	return model.TokenCookiesJwt{}
}
