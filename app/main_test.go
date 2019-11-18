package main

// Test Use Case not for Provider, please Remember

// func TestRedisGet(T *testing.T) {
// 	poolRedis := &redis.Pool{
// 		// Maximum number of idle connections in the pool.
// 		MaxIdle: 80,
// 		// max number of connections
// 		MaxActive: 12000,
// 		// Dial is an application supplied function for creating and
// 		// configuring a connection.
// 		Dial: func() (redis.Conn, error) {
// 			c, err := redis.Dial("tcp", ":6379")
// 			if err != nil {
// 				panic(err.Error())
// 			}
// 			return c, err
// 		},
// 	}
// 	connRedis := poolRedis.Get()

// 	value, err := redis.String(connRedis.Do("GET", "Artajasa.BPAYinfoBills"))
// 	if err != nil {
// 		T.Errorf("Error lah  %+v\n", err)
// 	}
// 	fmt.Printf("val redis %v", value)
// }
