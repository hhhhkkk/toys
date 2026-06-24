package entity

type User struct {
	Id      uint   `gorm:"primaryKey;autoIncrement"`
	Email    string `gorm:"size:255;uniqueIndex"`
	Password string `gorm:"size:255"`
	Status   int    `gorm:"type:tinyint(11);not null"`
}

func NewUser(id uint, email, password string, status int) *User {
	return &User{
		Id:       id,
		Email:    email,
		Password: password,
		Status:   status,
	}
}
