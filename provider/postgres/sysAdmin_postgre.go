package postgres

import (
	"AuthorizationJWT/engine"
	"AuthorizationJWT/model"

	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
)

type (
	sysAdminRepository struct {
		session *gorm.DB
	}
)

func newSysAdminRepostiory(db *gorm.DB) engine.SysAdminRepository {
	return &sysAdminRepository{db}
}

func (s *sysAdminRepository) Select(m *model.SysAdmin) (*model.SysAdmin, error) {
	result := s.session.Where("name = ?", m.Name).First(&m)
	if result.Error != nil {
		return m, errors.Errorf("Error Select SysAdmin : %v", result.Error)
	}
	return m, nil
}
