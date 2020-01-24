package postgres

import (
	"AuthorizationJWT/engine"
	"AuthorizationJWT/model"

	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
)

type (
	usersRepository struct {
		session *gorm.DB
	}
)

func newUsersRepository(db *gorm.DB) engine.UsersRepository {
	return &usersRepository{db}
}

func (s usersRepository) Insert(k *model.Users) error {
	result := s.session.Create(&k)
	if result.Error != nil {
		return errors.Errorf("Error Insert Users : %v", result.Error)
	}
	return nil
}

// select a row of users by username or id
func (s usersRepository) Select(k *model.Users) (*model.Users, error) {
	if k.UserName != "" {
		result := s.session.Where("user_name = ?", k.UserName).First(&k)
		if result.Error != nil {
			return k, errors.Errorf("Error Select Users : %v", result.Error)
		}
		return k, nil
	} else if k.ID != 0 {
		result := s.session.Where("id = ?", k.ID).First(&k)
		if result.Error != nil {
			return k, errors.Errorf("Error Select Users : %v", result.Error)
		}
		return k, nil
	}
	return &model.Users{}, nil
}

func (s usersRepository) UpdateAll(m *model.Users) error {
	// update map
	maps := map[string]interface{}{
		"password":       m.Password,
		"user_full_name": m.UserFullName,
		"email":          m.Email,
		"role_id":        m.RoleID,
		"activated":      m.Active,
	}
	result := s.session.Model(&m).Where("user_name = ?", m.UserName).Update(maps)
	if result.Error != nil {
		return errors.Errorf("Error Update Users : %v", result.Error)
	}
	return nil
}

func (s usersRepository) Remove(m *model.Users) error {
	result := s.session.Where("KEY = ?", m.UserName).Delete(&m)
	if result.Error != nil {
		return errors.Errorf("Error Delete a Users : %v", result.Error)
	}
	return nil
}
