package data

import (
	"context"
	"time"

	"github.com/hhhhkkk/mini-blog/internal/biz/admin/entity"
	"gorm.io/gorm"
)

type AdminUser struct {
	Id        int       `gorm:"primaryKey"`
	Name      string    `gorm:"type:varchar(255);not null"`
	Email     string    `gorm:"type:varchar(255);not null"`
	Password  string    `gorm:"type:varchar(255);not null"`
	Status    int       `gorm:"type:int(11);not null"`
	CreatedAt time.Time `gorm:"type:datetime;not null"`
	UpdatedAt time.Time `gorm:"type:datetime;not null"`
}

type AdminUserRepoImpl struct {
	db *gorm.DB
}

func (AdminUser) TableName() string {
	return "admin_users"
}

func NewAdminUserRepoImpl(db *gorm.DB) *AdminUserRepoImpl {
	return &AdminUserRepoImpl{
		db: db,
	}
}

func NewRepoImpl() *AdminUserRepoImpl {
	return &AdminUserRepoImpl{}
}

func (impl *AdminUserRepoImpl) EmailExist(email string, uid uint) (bool, error) {
	ctx := context.Background()
	db := gorm.G[AdminUser](impl.db)
	query := db.Where("email = ?", email)
	if uid > 0 {
		query = query.Where("id != ?", uid)
	}
	_, err := query.First(ctx)
	return err == nil, err
}

func (impl *AdminUserRepoImpl) CreateUser(user *entity.User) (*entity.User, error) {
	data := impl.db.Where(AdminUser{Email: user.Email}).FirstOrCreate(&user)
	if data.Error != nil {
		return nil, data.Error
	}
	return user, nil
}
