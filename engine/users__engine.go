package engine

import (
	"crypto/sha1"
)

type (
	// Users is the interface for interactor and this is Use Case
	Users interface {
		AddUsers(h *AddUserReq) *UsersDefaultResp
		UpdateUsers(h *UpdtUserReq) *UsersDefaultResp
		DeleteUsers(h *SelectUserReq) *UsersDefaultResp
		SelectUsers(h *SelectUserReq) *SelectUserResp
	}

	users struct {
		repository UsersRepository
	}

	// UsersDefaultResp for default Respon after success Respon
	UsersDefaultResp struct {
		ID    string
		Error string
	}
)

var (
	sha = sha1.New()
	key = "DOTA"
)

func (f *engineFactory) NewUsersEngine() Users {
	repostruc := &users{
		repository: f.NewUsersRespository(),
	}

	repostruc.InitializationUsers()
	return repostruc
}

// InitializationUsers create new default users administrator if not exist
func (s users) InitializationUsers() {
	// create new default users administrator if not exist

	// Default Admin
	UserName := "adminSkywalker"
	Password := "jaquest"
	UserFullName := "Dota MMR nya cuma 4k, Ancient"
	RoleID := 1
	Email := "dimashasbi.habibi@gmail.com"

	firstUser := SelectUserReq{
		ID:       "0",
		UserName: &UserName,
	}
	val := s.SelectUsers(&firstUser)
	if val.Error != "" || !val.Active && val.UserName == UserName {
		// no users = no data record = error
		// no rows = val.active is false and role ID is 0

		inp := &AddUserReq{
			ID:           "0",
			UserName:     firstUser.UserName,
			Password:     &Password,
			UserFullname: &UserFullName,
			RoleID:       &RoleID,
			Email:        &Email,
		}
		sil := s.AddUsers(inp)
		if sil.Error != "" {
			panic("Error input First users")
		}
	}
}
