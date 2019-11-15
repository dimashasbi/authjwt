package model

import (
	"time"
)

//PayloadCookiesJWT is used to represent payload in JWT
type PayloadCookiesJWT struct {
	Iss string `json:"iss"`
	Sub string `json:"sub"`
	Aud string `json:"aud"`
	Exp string `json:"exp"`
	Nbf string `json:"nbf"`
	Iat string `json:"iat"`
	Jti string `json:"jti"`
}

//VerifyExpireAt check expire token time
func (claim *PayloadCookiesJWT) VerifyExpireAt() (bool, error) {
	var (
		yyyymmddhhmmss = "2006-01-02 15:04:05"
	)
	durationExp, err := time.ParseDuration(claim.Exp)
	if err != nil {
		return false, err
	}
	issuedAt, _ := time.ParseInLocation(yyyymmddhhmmss, claim.Iat, time.Local)
	expireAt := issuedAt.Add(durationExp)
	now := time.Now()
	return expireAt.After(now), nil
}

//TokenCookiesJwt is used to represent data cookies token
type TokenCookiesJwt struct {
	HeaderPlusPayload string
	Signature         string
	RefreshToken      string
}
