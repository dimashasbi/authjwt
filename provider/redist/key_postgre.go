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
func (r *keyRepository) StoreToken(userData model.Users, jwtIDAccess string, jwtIDReresh string) error {
	keys := "tokenAuth:" + string(userData.ID) + ":" + jwtIDAccess
	_, err := r.redisConn.Do("SET", keys, jwtIDReresh)
	if err != nil {
		return err
	}
	return nil
}

//GetToken by idJWT with uuid format, search on redis
func (r *keyRepository) GetToken(userID string, jwtIDAccess string) (string, error) {
	key := "tokenAuth:" + userID + ":" + jwtIDAccess
	jwtIDReresh, err := redis.String(r.redisConn.Do("GET", key))
	if err != nil {
		return "", err
	}
	return jwtIDReresh, nil
}

//RemoveToken from redis
func (r *keyRepository) RemoveToken(userID, jwtIDAccess string) error {
	key := "tokenAuth:" + userID + ":" + jwtIDAccess
	_, err := r.redisConn.Do("DEL", key)
	if err != nil {
		return err
	}
	return nil
}
