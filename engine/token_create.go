package engine

import (
	"AuthorizationJWT/model"
	"fmt"
	"strconv"
	"strings"
	"time"
)

type (
	//CreateTokenReq input
	CreateTokenReq struct {
		userName string
		ID       int
	}

	//CreateTokenResp output
	CreateTokenResp struct {
		HeaderPlusPayload string
		Signature         string
		RefreshToken      string
		ErrorInfo         string
	}
)

func (t *token) CreateTokenUsecase(userData *CreateTokenReq) *CreateTokenResp {
	var (
		timeFormat = "2006-01-02 15:04:05"
	)
	timeNow := time.Now().Format(timeFormat)

	// get from database (just id and admin)

	usermod := t.getUserFromDB(userData)
	if !usermod.Active {
		return &CreateTokenResp{
			ErrorInfo: "10|Error no User to create Token",
		}
	}

	jwtIDAcc, _ := t.newUUID()
	jwtIDRfr, _ := t.newUUID()

	// create Access and Refresh Token
	subject := strconv.Itoa(usermod.ID) + "." + usermod.UserName
	payloadAccessJwt := t.mapper.ToPayloadJwt(subject, "2m", timeNow, jwtIDAcc)
	payloadRfrshJwt := t.mapper.ToPayloadJwt(subject, "30m", timeNow, jwtIDRfr)

	AccToken, RefrToken, _ := t.createToken(payloadAccessJwt, payloadRfrshJwt)
	AccTokenArr := strings.Split(AccToken, ".")

	// save to redis
	err := t.key.StoreToken(usermod.ID, payloadAccessJwt.JwtID, payloadRfrshJwt.JwtID)
	if err != nil {
		fmt.Printf("Error Store Token to Redis : %v+", err)
		return &CreateTokenResp{
			ErrorInfo: "11|Error Store Token",
		}
	}

	return &CreateTokenResp{
		HeaderPlusPayload: AccTokenArr[0] + "." + AccTokenArr[1],
		Signature:         AccTokenArr[2],
		RefreshToken:      RefrToken,
	}
}

func (t *token) getUserFromDB(userData *CreateTokenReq) *model.Users {
	mod := model.NewUsers(userData.userName, "", "", "", userData.ID, 0, 0, time.Time{}, false)
	var (
		result *model.Users
		err    error
	)

	if mod.UserName != "" {
		result, err = t.user.Select(mod)
		if err != nil {
			fmt.Printf("%+v", err)
			return result
		}
	} else if mod.ID != 0 {
		result, err = t.user.Select(mod)
		if err != nil {
			fmt.Printf("%+v", err)
			return result
		}
	}
	return result
}
