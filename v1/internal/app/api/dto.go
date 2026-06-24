package api

import "errors"

type RegisterReq struct {
	Email      string `json:"email"`
	Password   string `json:"password"`
	InviteCode string `json:"invite_code,omitempty"`
}

func (req *RegisterReq) Validate() error {
	if len(req.Email) <= 0 {
		return errors.New("username is required")
	}
	if len(req.Password) <= 0 {
		return errors.New("password is required")
	}
	return nil
}
