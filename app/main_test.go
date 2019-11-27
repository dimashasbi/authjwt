package main

import (
	"AuthorizationJWT/adapter"
	"AuthorizationJWT/engine"
	"AuthorizationJWT/model"
	"AuthorizationJWT/provider/fileconfig"
	"AuthorizationJWT/provider/postgres"
	"AuthorizationJWT/provider/redist"
	"testing"

	"github.com/gomodule/redigo/redis"
	"github.com/stretchr/testify/assert"
)

// TestFactory to Get All Structure for Test Case
type (
	TestFactory struct {
		dbconf     model.DBConfigurationModel
		eng        engine.EnginesFactory
		dbhand     engine.StorageFactory
		rdshand    engine.RedisFactory
		testEngine engine.TestingEngineStr
		adapter    adapter.Handler
	}
)

func (t *TestFactory) initializeApp() {
	// Initialize Application First
	t.dbconf = fileconfig.GetDBConfig()
	// Connect and Migrate DB
	t.dbhand = postgres.NewStorage(t.dbconf)
	// Connect to Redis
	t.rdshand = redist.NewRedis()
	// Prepare Engine for Use Case Logic
	t.eng = engine.NewEngine(t.dbhand, t.rdshand)
}

func TestSetKeytoRedis(t *testing.T) {
	userid := "hasbi"
	idTokenAccess := "ilmlzlaksjd"
	Expected := map[string]string{
		userid: idTokenAccess,
	}
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

	keys := "tokenAuth:" + userid + ":"
	_, err := connRedis.Do("SET", keys, idTokenAccess)
	if err != nil {
		t.Errorf(" Error Set key %v+ ", err)
	}

	// check actual
	value, err := redis.String(connRedis.Do("GET", keys))
	if err != nil {
		t.Errorf(" Error Get key %v+ ", err)
	}
	Actual := map[string]string{
		userid: value,
	}

	assert.Equal(t, Expected, Actual)
}

func TestGetToken(t *testing.T) {
	var (
		// TestEngine engine.TestingEngine
		TestStr engine.TestingEngineStr
	)

	testMain := &TestFactory{}
	testMain.initializeApp()

	TestStr = testMain.eng.NewTestEngine()

	expected := "HAHA"

	userID := ""
	ploadSign := ""
	actual, _ := TestStr.Key.GetToken(userID, ploadSign)

	assert.Equal(t, expected, actual)
}
