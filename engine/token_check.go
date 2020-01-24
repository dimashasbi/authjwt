package engine

import (
	"AuthorizationJWT/model"
	"fmt"
	"strconv"
	"strings"
)

type (
	// CheckTokenReq for input CheckToken
	CheckTokenReq struct {
		HeaderPlusPayload string
		Signature         string
		RefreshToken      string
	}

	// CheckTokenResp for input CheckToken
	CheckTokenResp struct {
		Valid       bool
		InvalidInfo string
		NewRfrToken string
	}
)

func (t *token) CheckTokenUsecase(c *CheckTokenReq) *CheckTokenResp {
	var (
		payloadAccToken, payloadRfrToken model.PayloadCookiesJWT
	)
	accToken := c.HeaderPlusPayload + "." + c.Signature
	rfrToken := c.RefreshToken

	// parse token first
	AccTokenParsed, RfrTokenParsed, err := t.parseToken(accToken, rfrToken)
	if err != nil {
		return &CheckTokenResp{
			Valid:       false,
			InvalidInfo: "20|Error Parsing",
		}
	} else if !AccTokenParsed.Valid || !RfrTokenParsed.Valid {
		return &CheckTokenResp{
			Valid:       false,
			InvalidInfo: "21|Invalid Parsing Token",
		}
	}

	// get payload
	payloadAccToken, payloadRfrToken, err = t.mapper.MarshalToPayloadCookiesJWT(AccTokenParsed, RfrTokenParsed)
	if err != nil {
		return &CheckTokenResp{
			Valid:       false,
			InvalidInfo: "22|Mapper Error",
		}
	}

	// check token in Redis
	arrSubAcc := strings.Split(payloadAccToken.Subject, ".")
	jtiTokenRefr, err := t.key.GetToken(arrSubAcc[0], payloadAccToken.JwtID)
	if err != nil {
		return &CheckTokenResp{
			Valid:       false,
			InvalidInfo: "23|Redis Error",
		}
	} else if jtiTokenRefr == "" {
		return &CheckTokenResp{
			Valid:       false,
			InvalidInfo: "24|Token not found, Do Generate One",
		}
	} else if payloadRfrToken.JwtID != jtiTokenRefr {
		return &CheckTokenResp{
			Valid:       false,
			InvalidInfo: "24|Token Keychain Refresh not Valid",
		}
	}

	// check alive token
	isAliveAccToken, err := payloadAccToken.VerifyExpireAt()
	if err != nil {
		fmt.Printf("Error check Expire : %+v", err)
	}
	isAliveRfrToken, err := payloadRfrToken.VerifyExpireAt()
	if err != nil {
		fmt.Printf("Error check Expire : %+v", err)
	}

	if isAliveAccToken == false && isAliveRfrToken == true {
		// create new Refresh Token
		// remove first
		t.key.RemoveToken(arrSubAcc[0], payloadAccToken.JwtID)
		// create new
		id, _ := strconv.Atoi(arrSubAcc[0])
		data := &CreateTokenReq{
			ID: id,
		}
		newTokenPayloadJWT := t.CreateTokenUsecase(data)
		return &CheckTokenResp{
			Valid:       true,
			NewRfrToken: newTokenPayloadJWT.RefreshToken,
		}

	} else if isAliveAccToken == false {
		// remove token,
		t.key.RemoveToken(arrSubAcc[0], payloadAccToken.JwtID)
	}

	return &CheckTokenResp{
		Valid: true,
	}
}
