package engine

import "AuthorizationJWT/model"

func (t *token) CheckTokenUsecase(m model.TokenCookiesJwt) interface{} {
	var (
		token = m.HeaderPlusPayload + "." + m.Signature
	)
	return model.Properties{}
}

func (t *token) CreateTokenUsecase(userData model.Users) model.TokenCookiesJwt {
	return model.TokenCookiesJwt{}
}
