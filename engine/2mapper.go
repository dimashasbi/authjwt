package engine

import (
	"AuthorizationJWT/model"
	"encoding/json"

	"github.com/dgrijalva/jwt-go"
)

type (
	// Mapper defines  methods that need to implement
	Mapper interface {
		ToJwtClaim(payload model.PayloadCookiesJWT) (jwt.MapClaims, error)
		ToPayloadJwt(subject, expired, createdtime, jwtid string) model.PayloadCookiesJWT
		MarshalToPayloadCookiesJWT(tokenParse *jwt.Token, refreshTokenParse *jwt.Token) (model.PayloadCookiesJWT, model.PayloadCookiesJWT, error)
	}

	mapper struct {
	}
)

func (f *engineFactory) NewMapper() Mapper {
	return &mapper{}
}

//ToJwtClaim is used to move from payloa to jwt.MapClaims on lib jwt `gdrijalva`
func (m *mapper) ToJwtClaim(payload model.PayloadCookiesJWT) (jwt.MapClaims, error) {
	var cPayload jwt.MapClaims
	asJSON, err := json.Marshal(payload)
	err = json.Unmarshal(asJSON, &cPayload)
	if err != nil {
		return nil, err
	}
	return cPayload, nil
}

//ToPayloadJwt is to mapping string to model PayloadCookiesJWT
func (m *mapper) ToPayloadJwt(subject, expired, createdtime, jwtid string) model.PayloadCookiesJWT {
	return model.PayloadCookiesJWT{
		Subject:     subject,
		Expired:     expired,
		CreatedTime: createdtime,
		JwtID:       jwtid,
	}
}

//MarshalToPayloadCookiesJWT the result is 1 payloadTokenOrdinary, 2 payload tokenRefresh, 3 error
func (m *mapper) MarshalToPayloadCookiesJWT(tokenParse *jwt.Token, refreshTokenParse *jwt.Token) (model.PayloadCookiesJWT, model.PayloadCookiesJWT, error) {
	var (
		payloadToken, payloadRefreshToken model.PayloadCookiesJWT
	)
	asJSON1, err := json.Marshal(tokenParse.Claims)
	if err != nil {
		return payloadToken, payloadRefreshToken, err
	}
	json.Unmarshal(asJSON1, &payloadToken)
	asJSON2, err := json.Marshal(refreshTokenParse.Claims)
	if err != nil {
		return payloadToken, payloadRefreshToken, err
	}
	json.Unmarshal(asJSON2, &payloadRefreshToken)
	return payloadToken, payloadRefreshToken, nil
}
