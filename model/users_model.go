package model

import (
	"time"
)

// Users is model for database configuration system
type Users struct {
	ID           int32  `gorm:"size:30;index:users_id;unique"`
	UserName     string `gorm:"size:30;index:users_name;unique"`
	Password     string `gorm:"size:50;index:users_pass"`
	UserFullName string `gorm:"size:50"`
	Email        string `gorm:"size:50"`
	RoleID       int    `gorm:"size:1;index:users_roleid"`
	// LastLoginTime      time.Time
	LoginFailCount int `gorm:"size:1;index:users_lgnfailcnt"`
	// LastChangePassTime time.Time
	// MustChangePass     bool `gorm:"index:users_muschgpas"`
	// DisableReasonCode  int  `gorm:"size:1;index:users_disbresncd"`
	CreatedTime time.Time
	// Disabled           bool `gorm:"index:users_disbl"`
	Active bool `gorm:"index:users_actv"`
}

// NewUsers create new Users Object
func NewUsers(username, password, userfullname, email string,
	role, loginfailcnt int, createdtime time.Time, active bool) *Users {
	return &Users{
		UserName:       username,
		Password:       password,
		UserFullName:   userfullname,
		Email:          email,
		RoleID:         role,
		LoginFailCount: loginfailcnt,
		CreatedTime:    createdtime,
		Active:         active,
	}
}
