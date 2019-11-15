package engine

import (
	"AccountManagement/model"
	"encoding/base64"
	"fmt"
	"time"
)

type (
	// UpdtUserReq for Request Update users
	UpdtUserReq struct {
		ID           string
		UserName     *string
		Password     *string
		UserFullname *string
		Email        *string
	}
)

// UpdateUsers use for change value of password, userfullname, email,
func (s *users) UpdateUsers(h *UpdtUserReq) *UsersDefaultResp {
	check := s.checkTagupdUser(h)
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
		0,
		0, // login fail count
		time.Time{},
		true, // active
	)
	err := s.repository.UpdateAll(usermod)
	if err != nil {
		fmt.Printf("%+v", err)

		return &UsersDefaultResp{
			ID:    string(h.ID),
			Error: "Error Update Users Table",
		}
	}

	return &UsersDefaultResp{
		ID:    h.ID,
		Error: "",
	}
}

func (s *users) checkTagupdUser(h *UpdtUserReq) string {

	if h.UserName == nil || *h.UserName == "" {
		return "Tag UserName is missing or empty "
	}

	if h.Password == nil || *h.Password == "" {
		return "Tag Password is missing or empty "
	}

	if h.UserFullname == nil {
		return "Tag Password is missing "
	}

	if h.Email == nil {
		return "Tag Email is missing "
	}
	return ""
}
