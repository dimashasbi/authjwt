package engine

import (
	"AccountManagement/model"
	"fmt"
	"strings"
	"time"
)

func (s users) DeleteUsers(h *SelectUserReq) *UsersDefaultResp {

	check := s.checkTagselUser(h)
	if check != "" {
		return &UsersDefaultResp{
			ID:    h.ID,
			Error: check,
		}
	}

	usermod := model.NewUsers(
		*h.UserName,
		"",
		"",
		"",
		0,
		0, // login fail count
		time.Time{},
		false, // active
	)
	err := s.repository.Remove(usermod)
	if err != nil {
		fmt.Printf("%+v", err)
		if strings.ContainsAny("record not found", err.Error()) {
			return &UsersDefaultResp{
				ID:    string(h.ID),
				Error: "Error No Data Recorded in Table Users",
			}
		}
		return &UsersDefaultResp{
			ID:    h.ID,
			Error: "Error Select to Table Users",
		}
	}
	return &UsersDefaultResp{
		ID:    h.ID,
		Error: "",
	}
}
