package redis

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

func newRedisRepository(rd redis.Conn) engine.KeyRepository {
	return &keyRepository{rd}
}

//StoreToken is used to store idToken & idTokenRefresh in redis
func (r *keyRepository) StoreToken(userData model.Users, idToken string, idTokenRefresh string) error {
	// keysFinder := "token:" + userData.UserName + ":*" //`token:admin:`
	// arrKeys, err := r.redisConn.Keys(keysFinder)
	// for _, key := range arrKeys {
	// 	r.redisEngine.Del(key)
	// }
	// keys := "token:" + userData.UserName + ":" + idToken //`token:admin:{uid}`
	// err = r.redisConn.Set(keys, idTokenRefresh, time.Minute*60)
	// if err != nil {
	// 	return err
	// }
	return nil
}

//GetToken by idJTI with uuid format, search on redis
func (r *keyRepository) GetToken(userID string, ploadSign string) (string, error) {
	// getkey := "token:" + userID + ":" + ploadSign
	// value, err := redis.String(r.redisConn.Do("GET", getkey))
	value, err := redis.String(r.redisConn.Do("GET", "Artajasa.BPAYinfoBills"))
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
