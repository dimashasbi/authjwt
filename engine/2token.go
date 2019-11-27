package engine

import (
	"AuthorizationJWT/model"
	"crypto/rsa"
)

type (
	// Token is the interface for interactor
	Token interface {
		Initialization()
		CheckTokenUsecase(ploadSign model.TokenCookiesJwt) interface{}
		CreateTokenUsecase(userData model.Users) model.TokenCookiesJwt
	}

	token struct {
		key        KeyRepository
		user       UsersRepository
		mapper     Mapper
		privKeyAcc *rsa.PrivateKey
		pubKeyAcc  *rsa.PublicKey
		privKeyRfr *rsa.PrivateKey
		pubKeyRfr  *rsa.PublicKey
	}
)

func (f *engineFactory) NewTokenEngines() Token {
	return &token{
		key:    f.NewKeyRepository(),
		user:   f.NewUsersRepository(),
		mapper: f.NewMapper(),
	}
}

func (t *token) Initialization() {

}
