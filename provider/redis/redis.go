package redis

import (
	"AuthorizationJWT/engine"
	"fmt"

	"github.com/gomodule/redigo/redis"
)

type (
	//redisFactory is to define struct of redis Jobs
	redisFactory struct {
		// RedisPool *redis.Pool
		RedisConn redis.Conn
	}
)

//NewRedis is to initalize redisFactory
func NewRedis() engine.RedisFactory {
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

	return &redisFactory{connRedis}
	// return &connRedis
}

func (r *redisFactory) NewRedisRepository() engine.KeyRepository {
	return newRedisRepository(r.RedisConn)
}

// ping tests connectivity for redis (PONG should be returned)
func (r *redisFactory) ping(c redis.Conn) error {

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
