package redist

import (
	"AuthorizationJWT/engine"
	"AuthorizationJWT/model"

	"github.com/gomodule/redigo/redis"
)

type (
	keyRepository struct {
		redisConn redis.Conn
	}
)

func newKeyRepository(rd redis.Conn) engine.KeyRepository {
	return &keyRepository{rd}
}

//StoreToken is used to store idToken & idTokenRefresh in redis
func (r *keyRepository) StoreToken(userData model.Users, idTokenAccess string, idTokenRefresh string) error {
	keys := "tokenAuth:" + userData.UserName + ":*" + idTokenAccess
	_, err := r.redisConn.Do("SET", keys, idTokenRefresh)
	if err != nil {
		return err
	}
	return nil
}

//GetToken by idJTI with uuid format, search on redis
func (r *keyRepository) GetToken(userID string, ploadSign string) (string, error) {
	// getkey := "token:" + userID + ":" + ploadSign
	// value, err := redis.String(r.redisConn.Do("GET", getkey))
	value, err := redis.String(r.redisConn.Do("GET", "HASBI"))
	if err != nil {
		return "", err
	}
	return value, nil
}

//RemoveToken from redis
func (r *keyRepository) RemoveToken(idToken string) error {
	// err := r.redisEngine.Expire(("token:" + idToken), 0)
	// if err != nil {
	// 	return err
	// }
	return nil
}
