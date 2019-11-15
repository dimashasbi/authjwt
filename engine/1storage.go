package engine

import (
	"AuthorizationJWT/model"
)

type (
	// UsersRepository defines the methods that any
	// data storage provider needs to implement to get
	// and store Users
	UsersRepository interface {
		Select(k *model.Users) (*model.Users, error)
	}

	// SysAdminRepository defines the methods that any
	// data storage provider needs to implement to get
	// and store SysAdmin
	SysAdminRepository interface {
		Select()
	}

	// RedisRepository defines the methods that any
	// data storage provider needs to implement to get
	// and store Redis
	RedisRepository interface {
		StoreToken(userData model.Users, idToken string, idTokenRefresh string) error
		GetToken(userID string, uuidTokenAccess string) (string, error)
		RemoveToken(idToken string) error
	}

	// StorageFactory is the interface that a storage
	// provider needs to implement so that the engine can
	// request repository instances as it needs them
	StorageFactory interface {
		NewUsersRepository() UsersRepository
		NewSysAdminRepository() SysAdminRepository
		NewRedisRepository() RedisRepository
	}
)
