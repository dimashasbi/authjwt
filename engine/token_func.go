package engine

import (
	"AuthorizationJWT/model"
	"crypto/rand"
	"fmt"
	"io"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func (t *token) CheckTokenUsecase(m model.TokenCookiesJwt) interface{} {

	return model.Properties{}
}

func (t *token) CreateTokenUsecase(userData model.Users) model.TokenCookiesJwt {
	var (
		tokenCookies model.TokenCookiesJwt
		timeFormat   = "1994-06-17 18:02:30"
	)
	timeNow := time.Now().Format(timeFormat)
	jwtID, _ := t.newUUID()

	// get from database
	usermod := t.GetUserFromDB(&userData)

	// create Access and Refresh Token
	subject := string(userData.ID) + "." + userData.UserName
	payloadAccessJwt := t.mapper.ToPayloadJwt(subject, "2m", timeNow, jwtID)
	payloadRfrshJwt := t.mapper.ToPayloadJwt(subject, "2m", timeNow, jwtID)

	AccToken, RefrToken, _ := t.createToken(payloadAccessJwt, payloadRfrshJwt)
	AccTokenArr := strings.Split(AccToken, ".")
	tokenCookies = model.TokenCookiesJwt{
		HeaderPlusPayload: AccTokenArr[0] + "." + AccTokenArr[1],
		Signature:         AccTokenArr[2],
		RefreshToken:      RefrToken,
	}

	// save to redis
	err := t.key.StoreToken(*usermod, payloadAccessJwt.JwtID, payloadRfrshJwt.JwtID)
	if err != nil {
		fmt.Printf("Error Store Token to Redis : %v+", err)
	}

	return tokenCookies
}

func (t *token) GetUserFromDB(userData *model.Users) *model.Users {
	mod := model.NewUsers(userData.UserName, "", "", "", 0, 0, time.Time{}, false)
	result, err := t.user.Select(mod)
	if err != nil {
		fmt.Printf("%+v", err)
		return userData
	}
	return result
}

//newUUID generates a random UUID according to RFC 4122
func (t *token) newUUID() (string, error) {
	uuid := make([]byte, 16)
	n, err := io.ReadFull(rand.Reader, uuid)
	if n != len(uuid) || err != nil {
		return "", err
	}
	// variant bits; see section 4.1.1
	uuid[8] = uuid[8]&^0xc0 | 0x80
	// version 4 (pseudo-random); see section 4.1.3
	uuid[6] = uuid[6]&^0xf0 | 0x40
	return fmt.Sprintf("%x-%x-%x-%x-%x", uuid[0:4], uuid[4:6], uuid[6:8], uuid[8:10], uuid[10:]), nil
}

func (t *token) createToken(payloadAcc, payloadRfrs model.PayloadCookiesJWT) (string, string, error) {
	MapClmPayloadAcc, _ := t.mapper.ToJwtClaim(payloadAcc)
	MapClmPayloadRfrs, _ := t.mapper.ToJwtClaim(payloadRfrs)

	algMethodAcc := jwt.GetSigningMethod("RS256")
	algMethodRfr := jwt.GetSigningMethod("RS256")

	AccToken := &jwt.Token{
		Header: map[string]interface{}{
			"typ": "JWT",
			"alg": algMethodAcc.Alg(),
			"kid": 1,
		},
		Claims: MapClmPayloadAcc,
		Method: algMethodAcc,
	}
	RfrToken := &jwt.Token{
		Header: map[string]interface{}{
			"typ": "JWT",
			"alg": algMethodRfr.Alg(),
			"kid": 1,
		},
		Claims: MapClmPayloadRfrs,
		Method: algMethodRfr,
	}

	AccessToken, err1 := AccToken.SignedString(t.privKeyAcc)
	if err1 != nil {
		fmt.Printf("Error Create Token %v+", err1)
	}
	RefreshToken, err2 := RfrToken.SignedString(t.privKeyRfr)
	if err2 != nil {
		fmt.Printf("Error Create Token %v+", err2)
	}
	return AccessToken, RefreshToken, nil
}
