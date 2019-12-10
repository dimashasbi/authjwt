package engine

import (
	"AuthorizationJWT/model"
	"bufio"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"os"
)

type (
	// Token is the interface for interactor
	Token interface {
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
	privAcc, pubAcc := getAccountKey()
	privRfr, pubRfr := getRefreshKey()
	return &token{
		key:        f.NewKeyRepository(),
		user:       f.NewUsersRepository(),
		mapper:     f.NewMapper(),
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
