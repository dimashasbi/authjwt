package postgres

import (
	"AuthorizationJWT/engine"

	"github.com/jinzhu/gorm"
)

type (
	sysAdminRepository struct {
		session *gorm.DB
	}
)

func newSysAdminRepostiory(db *gorm.DB) engine.UsersRepository {
	return &sysAdminRepository{db}
}


