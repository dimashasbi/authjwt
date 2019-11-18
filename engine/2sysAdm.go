package engine

import (
	"AuthorizationJWT/model"
)

type (
	// SysAdmin is the interface for interactor
	SysAdmin interface {
		Initialization()
		GetSysAdminValue(a *model.SysAdmin) *model.SysAdmin // get configuration from Database for Redis DB
	}

	sysAdmin struct {
		repository SysAdminRepository
	}
)

func (f *engineFactory) NewSysAdmin() SysAdmin {
	return &sysAdmin{
		repository: f.NewSysAdminRepository(),
	}
}

// Initialization use for check Table
func (s *sysAdmin) Initialization() {

}

// get JSON value in DB and return SysAdmin value
func (s *sysAdmin) GetSysAdminValue(a *model.SysAdmin) *model.SysAdmin {
	return &model.SysAdmin{}
}
