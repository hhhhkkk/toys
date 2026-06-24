package admin

import (
	"errors"
	"regexp"
)

type EmailExistsDTO struct {
	Email string `json:"email"`
	Uid   uint   `json:"uid,omitempty"`
}

type RegisterUserDTO struct {
	Email          string `json:"email"`
	Password       string `json:"password"`
	RepeatPassword string `json:"repeat_password"`
}

func (dto *EmailExistsDTO) Validate() error {
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
	if err := validatePassword(dto.Password); err != nil {
		return err
	}
	return nil
}

// validatePassword 校验密码强度
func validatePassword(password string) error {
	// 1. 检查长度 > 8
	if len(password) < 8 || len(password) >= 256 {
		return errors.New("password must be longer than 8 characters and less than 256 characters")
	}

	// 2. 检查是否包含大写字母
	if !regexp.MustCompile(`[A-Z]`).MatchString(password) {
		return errors.New("password must contain at least one uppercase letter")
	}

	// 2. 检查是否包含小写字母
	if !regexp.MustCompile(`[a-z]`).MatchString(password) {
		return errors.New("password must contain at least one lowercase letter")
	}

	// 2. 检查是否包含特殊符号
	if !regexp.MustCompile(`[!@#$%^&*(),.?":{}|<>]`).MatchString(password) {
		return errors.New("password must contain at least one special character")
	}

	// 2. 检查是否包含数字
	if !regexp.MustCompile(`[0-9]`).MatchString(password) {
		return errors.New("password must contain at least one digit")
	}

	// 3. 检查数字/字母连续重复不能超过3次
	if hasConsecutiveRepeats(password, 3) {
		return errors.New("password cannot contain more than 3 consecutive identical characters")
	}

	return nil
}

// hasConsecutiveRepeats 检查是否有连续重复字符
func hasConsecutiveRepeats(s string, maxConsecutive int) bool {
	if len(s) <= maxConsecutive {
		return false
	}
	runes := []rune(s)
	count := 1
	for i := 1; i < len(runes); i++ {
		if runes[i] == runes[i-1] {
			count++
			if count > maxConsecutive {
				return true
			}
		} else {
			count = 1
		}
	}
	return false
}
