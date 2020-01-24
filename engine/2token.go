package engine

import (
	"AuthorizationJWT/engine/mapperJWT"
	"AuthorizationJWT/model"
	"bufio"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"io"
	"os"

	"github.com/dgrijalva/jwt-go"
)

type (
	// Token is the interface for interactor
	Token interface {
		CheckTokenUsecase(c *CheckTokenReq) *CheckTokenResp
		CreateTokenUsecase(c *CreateTokenReq) *CreateTokenResp
	}

	token struct {
		key        KeyRepository
		user       UsersRepository
		mapper     mapperJWT.Mapper
		privKeyAcc *rsa.PrivateKey
		pubKeyAcc  *rsa.PublicKey
		privKeyRfr *rsa.PrivateKey
		pubKeyRfr  *rsa.PublicKey
	}
)

func (f *engineFactory) NewTokenEngines() Token {
	privAcc, pubAcc := getAccountKey()
	privRfr, pubRfr := getRefreshKey()
	return &token{
		key:        f.NewKeyRepository(),
		user:       f.NewUsersRepository(),
		mapper:     mapperJWT.NewMapper(),
		privKeyAcc: privAcc,
		pubKeyAcc:  pubAcc,
		privKeyRfr: privRfr,
		pubKeyRfr:  pubRfr,
	}
}

func getAccountKey() (*rsa.PrivateKey, *rsa.PublicKey) {
	privateKeyFile, err := os.Open("private-account-token.pem")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
		panic(err)
	}

	pemfileinfo, _ := privateKeyFile.Stat()
	var size int64
	size = pemfileinfo.Size()
	pembytes := make([]byte, size)
	buffer := bufio.NewReader(privateKeyFile)
	_, err = buffer.Read(pembytes)
	data, _ := pem.Decode([]byte(pembytes))
	privateKeyFile.Close()

	privateKeyImported, err := x509.ParsePKCS1PrivateKey(data.Bytes)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
		panic(err)
	}
	publicKeyRead := &privateKeyImported.PublicKey
	// fmt.Println("Private Key : ", privateKeyImported)
	// fmt.Println("Public key: ", publicKeyRead)
	return privateKeyImported, publicKeyRead
}

func getRefreshKey() (*rsa.PrivateKey, *rsa.PublicKey) {
	privateKeyFile, err := os.Open("private-refresh-token.pem")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
		panic(err)
	}

	pemfileinfo, _ := privateKeyFile.Stat()
	var size int64
	size = pemfileinfo.Size()
	pembytes := make([]byte, size)
	buffer := bufio.NewReader(privateKeyFile)
	_, err = buffer.Read(pembytes)
	data, _ := pem.Decode([]byte(pembytes))
	privateKeyFile.Close()

	privateKeyImported, err := x509.ParsePKCS1PrivateKey(data.Bytes)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
		panic(err)
	}
	publicKeyRead := &privateKeyImported.PublicKey
	// fmt.Println("Private Key : ", privateKeyImported)
	// fmt.Println("Public key: ", publicKeyRead)
	return privateKeyImported, publicKeyRead
}

//newUUID generates a random UUID according to RFC 4122 used for jwtid access
func (t *token) newUUID() (string, error) {
	uuid := make([]byte, 16)
	n, err := io.ReadFull(rand.Reader, uuid)
	if n != len(uuid) || err != nil {
		return "", err
	}
	// variant bits;
	uuid[8] = uuid[8]&^0xc0 | 0x80
	// version 4 (pseudo-random);
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
	// fmt.Printf("Token %v+", AccessToken)
	RefreshToken, err2 := RfrToken.SignedString(t.privKeyRfr)
	if err2 != nil {
		fmt.Printf("Error Create Token %v+", err2)
	}
	return AccessToken, RefreshToken, nil
}

func (t *token) parseToken(accToken string, rfrToken string) (*jwt.Token, *jwt.Token, error) {
	accTokenParse, err := jwt.Parse(accToken, func(token *jwt.Token) (interface{}, error) { return t.pubKeyAcc, nil })
	if err != nil {
		return nil, nil, err
	}
	refreshTokenParse, err := jwt.Parse(rfrToken, func(token *jwt.Token) (interface{}, error) { return t.pubKeyRfr, nil })
	if err != nil {
		return nil, nil, err
	}
	return accTokenParse, refreshTokenParse, nil
}
