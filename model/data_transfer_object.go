package model

import (
	"time"
)

//PayloadCookiesJWT is used to represent payload in JWT
type PayloadCookiesJWT struct {
	Subject     string `json:"sub"` // user data and roles
	Expired     string `json:"exp"` // how minutes token gonna be expired
	CreatedTime string `json:"crt"` // create Token at
	JwtID       string `json:"jti"` // JWT ID for authorization ID of JWT
}

//VerifyExpireAt check expire token time
func (claim *PayloadCookiesJWT) VerifyExpireAt() (bool, error) {
	var (
		yyyymmddhhmmss = "2006-01-02 15:04:05"
	)
	durationExp, err := time.ParseDuration(claim.Expired)
	if err != nil {
		return false, err 
	}
	issuedAt, _ := time.ParseInLocation(yyyymmddhhmmss, claim.CreatedTime, time.Local)

	expireAt := issuedAt.Add(durationExp)
	now := time.Now()
	return expireAt.After(now), nil
}

//TokenCookiesJwt is used to represent data cookies token in JSON, Send and Receive
// as input and Output Authorization Service
type TokenCookiesJwt struct {
	HeaderPlusPayload string
	Signature         string
	RefreshToken      string
}
