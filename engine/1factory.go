package engine

type (
	// EnginesFactory interface allows us to provide
	// other parts of the system with a way to make
	// instances of our use-case / interactors when
	// they need to
	EnginesFactory interface {
		NewTokenEngines() Token
		NewUsers() Users
		NewSysAdmin() SysAdmin
		NewTestEngine() TestingEngineStruct
	}

	engineFactory struct {
		StorageFactory
		RedisFactory
	}
)

// NewEngine creates a new engine factory that will
// make use of the passed in StorageFactory for any
// data persistence needs.
func NewEngine(s StorageFactory, r RedisFactory) EnginesFactory {
	return &engineFactory{s, r}
}
