package engine

type (
	// TestingEngineStruct Structure consist of Repository Interface.
	TestingEngineStruct struct {
		Key      KeyRepository
		User     UsersRepository
		SysAdmin SysAdminRepository
		Mapper   Mapper
	}
)

func (f *engineFactory) NewTestEngine() TestingEngineStruct {
	return TestingEngineStruct{
		Key:      f.NewKeyRepository(),
		User:     f.NewUsersRepository(),
		SysAdmin: f.NewSysAdminRepository(),
		Mapper:   f.NewMapper(),
	}
}
