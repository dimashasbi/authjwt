package engine

import (
	"AccountManagement/model"
	"fmt"
	"strings"
	"time"
)

type (
	// SelectUserReq for request select User
	SelectUserReq struct {
		ID       string
		UserName *string
	}

	// SelectUserResp for respond select User
	SelectUserResp struct {
		ID             string
		UserName       string
		UserFullName   string
		Email          string
		RoleID         int
		LoginFailCount int
		Active         bool
		Error          string
	}
)

func (s users) SelectUsers(h *SelectUserReq) *SelectUserResp {

	check := s.checkTagselUser(h)
	if check != "" {
		return &SelectUserResp{
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
	sult, err := s.repository.Select(usermod)
	if err != nil {
		fmt.Printf("%+v", err)
		if strings.ContainsAny("record not found", err.Error()) {
			return &SelectUserResp{
				ID:    string(h.ID),
				Error: "Error No Data Recorded in Table Users",
			}
		}
		return &SelectUserResp{
			ID:       h.ID,
			UserName: *h.UserName,
			Error:    "Error Select to Table Users",
		}
	}
	return &SelectUserResp{
		ID:             h.ID,
		UserName:       sult.UserName,
		UserFullName:   sult.UserFullName,
		Email:          sult.Email,
		RoleID:         sult.RoleID,
		LoginFailCount: sult.LoginFailCount,
		Active:         sult.Active,
		Error:          "",
	}
}

func (s *users) checkTagselUser(h *SelectUserReq) string {

	if h.UserName == nil || *h.UserName == "" {
		return "Tag UserName is missing or empty "
	}
	return ""
}
