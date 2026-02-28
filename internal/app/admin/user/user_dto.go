package user

import (
	"errors"
	"regexp"
)

type EmailExistsEmailDTO struct {
	Email string `json:"email"`
	Uid   uint   `json:"uid"`
}

type RegisterUserDTO struct {
	Email          string `json:"email"`
	Password       string `json:"password"`
	RepeatPassword string `json:"repeat_password"`
}

func (dto *EmailExistsEmailDTO) Validate() error {
	if dto.Uid <= 0 {
		return errors.New("uid is required")
	}
	if len(dto.Email) <= 0 {
		return errors.New("email is required")
	}
	// 校验邮箱格式, 用正则, 且邮箱长度不能超过 255 个字符（utf8mb4）
	if !regexp.MustCompile(`^[a-zA-Z0-9.!#$%&'*+/=?^_{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$`).MatchString(dto.Email) {
		return errors.New("email format is invalid")
	}
	if len(dto.Email) > 255 {
		return errors.New("email is too long")
	}
	return nil
}

func (dto *RegisterUserDTO) Validate() error {
	if dto.Password != dto.RepeatPassword {
		return errors.New("password and repeat_password are not equal")
	}
	if len(dto.Password) < 6 {
		return errors.New("password is too short")
	}
	if len(dto.Password) > 255 {
		return errors.New("password is too long")
	}
	// 检测密码安全强度, 必须还有大、小写字母、数字、特殊字符 !@#$%^&*, 且数字、字母连续次数不能超过 3
	if !regexp.MustCompile(`^(?=.*[a-z])(?=.*[A-Z])(?=.*[0-9])(?=.*[!@#$%^&*])(?=\S+$).{6,255}$`).MatchString(dto.Password) {
		return errors.New("password is not secure")
	}
	return nil
}
