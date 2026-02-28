package entity

type User struct {
	Uid      uint   `gorm:"primaryKey;autoIncrement"`
	Email    string `gorm:"size:255;uniqueIndex"`
	Password string `gorm:"size:255"`
}

func NewUser(uid uint, email string, password string) *User {
	return &User{
		Uid:      uid,
		Email:    email,
		Password: password,
	}
}
