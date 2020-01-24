package main

import (
	"AuthorizationJWT/adapter"
	"AuthorizationJWT/engine"
	"AuthorizationJWT/model"
	"AuthorizationJWT/provider/fileconfig"
	"AuthorizationJWT/provider/postgres"

	// "AuthorizationJWT/provider/postgres"
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
		testEngine engine.TestingEngineStruct
		adapter    adapter.Handler
	}
	redisStruct struct {
		redisConn redis.Conn
	}
)

var (
	usingDB    = true
	usingRedis = true
)

func (t *TestFactory) initializeApp() {
	// Initialize Application First
	t.dbconf = fileconfig.GetDBConfig()

	// Activate 3rd Party Application
	// Connect and Migrate DB
	t.dbhand = postgres.NewStorage(t.dbconf)
	// Connect to Redis
	t.rdshand = redist.NewRedis()

	// Prepare Engine for Use Case Logic
	t.eng = engine.NewEngine(t.dbhand, t.rdshand)
	// Set Struct for Testing
	t.testEngine = t.eng.NewTestEngine()

}

func (t *TestFactory) initializeRedis() {
	// Initialize Application First
	t.dbconf = fileconfig.GetDBConfig()

	// Activate 3rd Party Application
	// Connect to Redis
	t.rdshand = redist.NewRedis()

	// Prepare Engine for Use Case Logic
	t.eng = engine.NewEngine(t.dbhand, t.rdshand)
	// Set Struct for Testing
	t.testEngine = t.eng.NewTestEngine()

}

func TestSetKeytoRedis(t *testing.T) {
	userid := "hasbi"
	idTokenAccess := "ilmlzlaksjd"
	Expected := map[string]string{
		userid: idTokenAccess,
	}

	// init redis conn
	var rds redisStruct
	rds.initRedis()

	keys := "tokenAuth:" + userid + ":"
	_, err := rds.redisConn.Do("SET", keys, idTokenAccess)
	if err != nil {
		t.Errorf(" Error Set key %v+ ", err)
	}

	// check actual
	value, err := redis.String(rds.redisConn.Do("GET", keys))
	if err != nil {
		t.Errorf(" Error Get key %v+ ", err)
	}
	Actual := map[string]string{
		userid: value,
	}

	assert.Equal(t, Expected, Actual)
}

func (rdtest *redisStruct) initRedis() {
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
	rdtest.redisConn = poolRedis.Get()
}

// Test Redis Get Token
func TestGetToken(t *testing.T) {
	// basic delcaration
	var (
		TestStr engine.TestingEngineStruct
	)
	testMain := &TestFactory{}
	testMain.initializeRedis()
	TestStr = testMain.testEngine

	var rds redisStruct
	rds.initRedis()

	// Do Test
	JWTIDrefreshExpected := "HAMSSS"
	usermod := model.Users{
		ID: 1,
	}
	JWTIDaccess := "ADAKA"

	// Set Token direct to Redis
	key := "tokenAuth:" + string(usermod.ID) + ":" + JWTIDaccess
	_, err := rds.redisConn.Do("SET", key, JWTIDrefreshExpected)
	if err != nil {
		t.Errorf(" Error Set key %v+ ", err)
	}

	JWTIDrefreshActual, _ := TestStr.Key.GetToken(string(usermod.ID), JWTIDaccess)

	assert.Equal(t, JWTIDrefreshExpected, JWTIDrefreshActual)

	// clean environment
	_, err = rds.redisConn.Do("DEL", key)
}

// Test Redis Get Index Token ?
// Test Redis Store Token
func TestStoreToken(t *testing.T) {
	// basic delcaration
	var (
		TestStr engine.TestingEngineStruct
	)
	testMain := &TestFactory{}
	testMain.initializeRedis()
	TestStr = testMain.testEngine

	var rds redisStruct
	rds.initRedis()

	// Do Test
	usermod := model.Users{
		ID: 1,
	}
	JWTIDaccess := "dimashasbi"
	JWTIDrefreshExpected := "habibiiiii"

	err := TestStr.Key.StoreToken(usermod, JWTIDaccess, JWTIDrefreshExpected)
	if err != nil {
		t.Errorf("Error Store Token : %v+", err)
	}

	// Get Token direct to Redis
	key := "tokenAuth:" + string(usermod.ID) + ":" + JWTIDaccess
	JWTIDrefreshActual, err := redis.String(rds.redisConn.Do("GET", key))
	if err != nil {
		t.Errorf(" Error Get key %v+ ", err)
	}

	assert.Equal(t, JWTIDrefreshExpected, JWTIDrefreshActual)

	// clean environment
	_, err = rds.redisConn.Do("DEL", key)
}

// Test Redis Set Index Token ?
// Test Redis Remove Token
func TestRemoveToken(t *testing.T) {
	// basic delcaration
	var (
		TestStr engine.TestingEngineStruct
	)
	testMain := &TestFactory{}
	testMain.initializeRedis()
	TestStr = testMain.testEngine

	var rds redisStruct
	rds.initRedis()

	// Do Test
	usermod := model.Users{
		ID: 1,
	}
	JWTIDaccess := "dimashasbi"
	idTokenAccess := "hahah"

	// set value direct to Redis
	key := "tokenAuth:" + string(usermod.ID) + ":" + JWTIDaccess
	_, err := rds.redisConn.Do("SET", key, idTokenAccess)
	if err != nil {
		t.Errorf(" Error Set key %v+ ", err)
	}

	// Remove Token direct to Redis
	err = TestStr.Key.RemoveToken(string(usermod.ID), JWTIDaccess)
	if err != nil {
		t.Errorf(" Error Remove key %v+ ", err)
	}

	_, err = redis.String(rds.redisConn.Do("GET", key))

	valueExpected := "redigo: nil returned"
	assert.Equal(t, valueExpected, err.Error())

	// clean environment = no need
}

// Test Redis Delete An Token of An Index ?
// Test Redis Delete All Index ?

// Test Use Case Create Token
func TestCreateToken(t *testing.T) {
	// basic delcaration
	var (
		Token engine.Token
	)
	testMain := &TestFactory{}
	testMain.initializeRedis()
	// TestStr = testMain.testEngine
	Token = testMain.eng.NewTokenEngines()

	// Do the Test

	// trigger by hit URL input JSON
	inp := model.Users{
		ID:           3,
		UserName:     "jaquest",
		UserFullName: "Dimas Hasbi Habibi",
	}
	Actual := Token.CreateTokenUsecase(inp)
	Expected := model.TokenCookiesJwt{
		HeaderPlusPayload: "",
		Signature:         "",
		RefreshToken:      "",
	}

	// 1. Check output message json
	assert.Equal(t, Expected, Actual)

	// 2. Check Token that stored at Redis
	JWTIDaccess := ""
	// Get Token direct to Redis
	var rds redisStruct
	rds.initRedis()
	key := "tokenAuth:" + string(inp.ID) + ":" + JWTIDaccess
	JWTIDrefreshActual, err := redis.String(rds.redisConn.Do("GET", key))
	if err != nil {
		t.Errorf(" Error Get key %v+ ", err)
	}

	// check jwt id refresh
	JWTIDrefreshExpected := ""
	assert.Equal(t, JWTIDrefreshExpected, JWTIDrefreshActual)

	// clean environment

	// clean redis

}

// Test Use Case Check Token
func TestCheckToken(t *testing.T) {
	// basic delcaration
	var (
		Token engine.Token
	)
	testMain := &TestFactory{}
	testMain.initializeRedis()
	// TestStr = testMain.testEngine
	Token = testMain.eng.NewTokenEngines()

	// Do the Test

	// trigger by hit URL input JSON
	inp := model.TokenCookiesJwt{
		HeaderPlusPayload: "",
		Signature:         "",
		RefreshToken:      "",
	}
	Actual := Token.CheckTokenUsecase(inp)
	Expected := model.TokenCookiesJwt{
		HeaderPlusPayload: "",
		Signature:         "",
		RefreshToken:      "",
	}

	// what test
	assert.Equal(t, Expected, Actual)

	// clean environment

}
