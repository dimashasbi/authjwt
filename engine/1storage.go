package engine

import (
	"AccountManagement/model"
)

type (
	// UsersRepository defines the methods that any
	// data storage provider needs to implement to get
	// and store greetings
	UsersRepository interface {
		Insert(k *model.Users) error
		Select(k *model.Users) (*model.Users, error)
		UpdateAll(k *model.Users) error
		Remove(k *model.Users) error
	}

	// StorageFactory is the interface that a storage
	// provider needs to implement so that the engine can
	// request repository instances as it needs them
	StorageFactory interface {
		// NewUsersRepository returns a storage specific
		NewUsersRespository() UsersRepository
	}
)
