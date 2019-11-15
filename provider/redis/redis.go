package redis

import (
	"github.com/go-redis/cache"
	"time"
)


type (
	//redisFactory is to define struct of redis Jobs
	redisFactory struct {
		redisEngine cache.IRedis
		timeout     time.Duration
	}	
)

//NewRedisRepository is to initalize redisFactory
func NewRedisRepository(redisEngine cache.IRedis, timeout time.Duration) handler.IredisFactory {
	return &redisFactory{redisEngine, timeout}
}

//StoreToken is used to store idToken & idTokenRefresh in redis
func (r *redisFactory) StoreToken(userData model.User, idToken string, idTokenRefresh string) error {
	keysFinder := "token:" + userData.ID + ":*" //`token:adin:`
	arrKeys, err := r.redisEngine.Keys(keysFinder)
	for _, key := range arrKeys {
		r.redisEngine.Del(key)
	}
	keys := "token:" + userData.ID + ":" + idToken //`token:admin:{uid}`
	err = r.redisEngine.Set(keys, idTokenRefresh, time.Minute*60)
	if err != nil {
		return err
	}
	return nil
}

//GetToken by idJTI with uuid format, search on redis
func (r *redisFactory) GetToken(userID string, ploadSign string) (string, error) {
	value, err := r.redisEngine.Get("token:" + userID + ":" + ploadSign)
	if err != nil {
		return "", err
	}
	return value, nil
}

//RemoveToken from redis
func (r *redisFactory) RemoveToken(idToken string) error {
	err := r.redisEngine.Expire(("token:" + idToken), 0)
	if err != nil {
		return err
	}
	return nil
}
