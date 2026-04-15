package data

import (
	"time"

	"github.com/dromara/carbon/v2"
	"github.com/hhhhkkk/mini-blog/internal/biz/api/entity"
)

type User struct {
	Id           int       `gorm:"primaryKey"`
	Name         string    `gorm:"type:varchar(255);not null"`
	Email        string    `gorm:"type:varchar(255);not null"`
	Password     string    `gorm:"type:varchar(255);not null"`
	Age          int       `gorm:"type:int(11);not null;"`
	InviteCode   string    `gorm:"type:varchar(50);not null"`
	Status       int       `gorm:"type:int(11);not null"`
	CreatedAt    time.Time `gorm:"type:datetime;not null"`
	UpdatedAt    time.Time `gorm:"type:datetime;not null"`
	LastLoginAt  time.Time `gorm:"type:datetime;not null"`
	LastActiveAt time.Time `gorm:"type:datetime;not null"`
}

func (u *User) TableName() string {
	return "users"
}

type UserRepoImpl struct {
	db *DB
}

func NewUserRepoImpl(db *DB) *UserRepoImpl {
	return &UserRepoImpl{
		db: db,
	}
}

func (impl *UserRepoImpl) EmailExists(email string, id int) bool {
	query := impl.db.GetClient().Model(&User{}).Where("email = ?", email)
	if id > 0 {
		query.Where("id != ?", id)
	}
	var count int64
	query.Count(&count)
	return count > 0
}

func (impl *UserRepoImpl) GetUserByInviteCode(code string) (*entity.UserEntity, error) {
	var u User
	tx := impl.db.GetClient().Model(&User{}).Where("invite_code = ?", code).First(&u)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return &entity.UserEntity{
		Id:           u.Id,
		Name:         u.Name,
		Email:        u.Email,
		Age:          u.Age,
		InviteCode:   u.InviteCode,
		Status:       u.Status,
		CreatedAt:    u.CreatedAt,
		UpdatedAt:    u.UpdatedAt,
		LastLoginAt:  u.LastLoginAt,
		LastActiveAt: u.LastActiveAt,
	}, nil
}

func (impl *UserRepoImpl) CreateUser(en *entity.UserEntity) (*entity.UserEntity, error) {
	u := &User{
		Name:         en.Name,
		Email:        en.Email,
		Age:          en.Age,
		InviteCode:   en.InviteCode,
		Status:       en.Status,
		CreatedAt:    en.CreatedAt,
		UpdatedAt:    en.UpdatedAt,
		LastLoginAt:  carbon.Now().StdTime(),
		LastActiveAt: carbon.Now().StdTime(),
	}
	tx := impl.db.GetClient().Model(&User{}).Create(u)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return &entity.UserEntity{
		Id:         u.Id,
		Name:       u.Name,
		Email:      u.Email,
		Age:        u.Age,
		InviteCode: u.InviteCode,
		Status:     u.Status,
	}, nil
}
