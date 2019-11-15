package engine

import (
	"AccountManagement/model"
	"encoding/base64"
	"fmt"
	"time"
)

type (
	// AddUserReq for Request Create users
	AddUserReq struct {
		ID           string
		UserName     *string
		Password     *string
		UserFullname *string
		RoleID       *int
		Email        *string
	}
)

func (s *users) AddUsers(h *AddUserReq) *UsersDefaultResp {
	check := s.checkTagaddUser(h)
	if check != "" {
		return &UsersDefaultResp{
			ID:    h.ID,
			Error: check,
		}
	}

	sha.Write([]byte(*h.UserName + "|" + *h.Password))
	encrPass := base64.StdEncoding.EncodeToString(sha.Sum([]byte(key)))

	usermod := model.NewUsers(
		*h.UserName,
		encrPass,
		*h.UserFullname,
		*h.Email,
		*h.RoleID,
		0, // login fail count
		time.Now(),
		true, // active
	)
	err := s.repository.Insert(usermod)
	if err != nil {
		fmt.Printf("%+v", err)

		return &UsersDefaultResp{
			ID:    h.ID,
			Error: "Error input to Users Table",
		}
	}

	return &UsersDefaultResp{
		ID:    h.ID,
		Error: "",
	}
}

func (s *users) checkTagaddUser(h *AddUserReq) string {

	if h.UserName == nil || *h.UserName == "" {
		return "Tag UserName is missing or empty "
	}

	if h.Password == nil || *h.Password == "" {
		return "Tag Password is missing or empty "
	}

	if h.UserFullname == nil {
		return "Tag Password is missing "
	}

	if h.RoleID == nil || *h.RoleID == 0 {
		return "Tag RoleID is missing or empty "
	}

	if h.Email == nil {
		return "Tag Email is missing "
	}

	return ""
}
