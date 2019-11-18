package engine

type (
	// Users is the interface for interactor and this is Use Case
	Users interface {
		SelectUsers(h *SelectUserReq) *SelectUserResp
	}

	users struct {
		repository UsersRepository
	}
)

func (f *engineFactory) NewUsers() Users {
	repostruc := &users{
		repository: f.NewUsersRepository(),
	}
	return repostruc
}
