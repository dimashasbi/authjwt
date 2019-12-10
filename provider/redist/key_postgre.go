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

// StoreToken is used to store idTokenAccess & idTokenRefresh in redis
func (r *keyRepository) StoreToken(userData model.Users, jwtIDAccess string, jwtIDReresh string) error {
	keys := "tokenAuth:" + string(userData.ID) + ":" + jwtIDAccess
	// set key
	_, err := r.redisConn.Do("SET", keys, jwtIDReresh)
	if err != nil {
		return err
	}
	// add into index
	keysIndex := "tokenAuth:" + string(userData.ID) + ":"
	_, err = r.redisConn.Do("SADD", keysIndex, jwtIDAccess)
	if err != nil {
		return err
	}
	return nil
}

// StoreToken is used to store idTokenAccess & idTokenRefresh in redis
func (r *keyRepository) AddIndex(userData model.Users, jwtIDAccess string, jwtIDReresh string) error {
	// add into index
	keysIndex := "tokenAuth:" + string(userData.ID) + ":"
	_, err := r.redisConn.Do("SADD", keysIndex, jwtIDAccess)
	if err != nil {
		return err
	}
	return nil
}

// GetToken by idJWT with uuid format, search on redis
func (r *keyRepository) GetToken(userID string, jwtIDAccess string) (string, error) {
	key := "tokenAuth:" + userID + ":" + jwtIDAccess
	jwtIDReresh, err1 := redis.String(r.redisConn.Do("GET", key))
	if err1 != nil {
		return "", err1
	}
	return jwtIDReresh, nil
}

// GetToken by idJWT with uuid format, search on redis
func (r *keyRepository) ListToken(userID string, jwtIDAccess string) ([]string, error) {
	keyind := "tokenAuth:" + userID + ":"
	keys, err := redis.String(r.redisConn.Do("SMEMBERS", keyind))
	var ListToken []string
	if err != nil {
		return ListToken, err
	}
	for _, key := range keys {
		ListToken = append(ListToken, string(key))
	}
	return ListToken, nil
}

// RemoveToken from redis
func (r *keyRepository) RemoveToken(userID, jwtIDAccess string) error {
	key := "tokenAuth:" + userID + ":" + jwtIDAccess
	// delete key
	_, err := r.redisConn.Do("DEL", key)
	if err != nil {
		return err
	}
	return nil
}

// RemoveToken from redis
func (r *keyRepository) RemoveAnIndex(userID, jwtIDAccess string) error {
	key := "tokenAuth:" + userID + ":"
	// delete key
	_, err := r.redisConn.Do("DEL", key)
	if err != nil {
		return err
	}
	return nil
}

// RemoveToken from redis
func (r *keyRepository) RemoveIndexToken(userID, jwtIDAccess string) error {
	// delete part of index
	keysIndexs := "tokenAuth:" + userID + ":"
	_, err := r.redisConn.Do("SREM", keysIndexs, jwtIDAccess)
	if err != nil {
		return err
	}
	return nil
}
