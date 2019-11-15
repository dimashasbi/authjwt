package adapter

import (
	"AccountManagement/engine"
	"encoding/json"
	"net/http"
)

type (
	users struct {
		engine.Users
	}
)

// AddUser for Add new User
func (a *users) AddUser(w http.ResponseWriter, r *http.Request) {
	mod := engine.AddUserReq{}
	json.NewDecoder(r.Body).Decode(&mod)

	resp := a.AddUsers(&mod)

	hasil, _ := json.Marshal(resp)
	w.Header().Set("Content-Type", "application/json")
	DefaultRespon(w, hasil)
	return
}

// SelectUser for Select User from database
func (a *users) SelectUser(w http.ResponseWriter, r *http.Request) {
	mod := engine.SelectUserReq{}
	json.NewDecoder(r.Body).Decode(&mod)

	resp := a.SelectUsers(&mod)

	hasil, _ := json.Marshal(resp)
	w.Header().Set("Content-Type", "application/json")
	DefaultRespon(w, hasil)
	return
}

// UpdateUsers for Update value user
func (a *users) UpdateUser(w http.ResponseWriter, r *http.Request) {
	mod := engine.UpdtUserReq{}
	json.NewDecoder(r.Body).Decode(&mod)

	resp := a.UpdateUsers(&mod)

	hasil, _ := json.Marshal(resp)
	w.Header().Set("Content-Type", "application/json")
	DefaultRespon(w, hasil)
	return
}
