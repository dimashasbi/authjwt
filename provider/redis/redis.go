package redis

import (
	"AuthorizationJWT/model"
	"fmt"

	"github.com/gomodule/redigo/redis"
)

type (
	//redisFactory is to define struct of redis Jobs
	redisFactory struct {
		redisPool *redis.Pool
		redisConn redis.Conn
	}
)

//NewRedisRepository is to initalize redisFactory
func NewRedisRepository() *redisFactory {
	poolRedis := &redis.Pool{
		// Maximum number of idle connections in the pool.
		MaxIdle: 80,
		// max number of connections
		MaxActive: 12000,
		// Dial is an application supplied function for creating and
		// configuring a connection.
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", ":6379")
			if err != nil {
				panic(err.Error())
			}
			return c, err
		},
	}
	connRedis := poolRedis.Get()

	return &redisFactory{
		poolRedis, connRedis,
	}
}

//StoreToken is used to store idToken & idTokenRefresh in redis
func (r *redisFactory) StoreToken(userData model.Users, idToken string, idTokenRefresh string) error {
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
func (r *redisFactory) GetToken(userID string, ploadSign string) (string, error) {
	getkey := "token:" + userID + ":" + ploadSign
	value, err := redis.String(r.redisConn.Do("GET", getkey))
	if err != nil {
		return "", err
	}
	return value, nil
}

//RemoveToken from redis
func (r *redisFactory) RemoveToken(idToken string) error {
	// err := r.redisEngine.Expire(("token:" + idToken), 0)
	// if err != nil {
	// 	return err
	// }
	return nil
}

// ping tests connectivity for redis (PONG should be returned)
func (r redisFactory) ping(c redis.Conn) error {

	// Send PING command to Redis
	pong, err := c.Do("PING")
	if err != nil {
		return err
	}

	// PING command returns a Redis "Simple String"
	// Use redis.String to convert the interface type to string
	s, err := redis.String(pong, err)
	if err != nil {
		return err
	}

	fmt.Printf("PING Response = %s\n", s)
	// Output: PONG

	return nil
}
