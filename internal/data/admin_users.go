package data

import (
	"time"

	"github.com/hhhhkkk/mini-blog/internal/biz/admin/entity"
)

type AdminUser struct {
	Id        int       `gorm:"primaryKey"`
	Name      string    `gorm:"type:varchar(255);not null"`
	Email     string    `gorm:"type:varchar(255);not null"`
	Password  string    `gorm:"type:varchar(255);not null"`
	Status    int       `gorm:"type:tinyint(11);not null"`
	CreatedAt time.Time `gorm:"type:datetime;not null"`
	UpdatedAt time.Time `gorm:"type:datetime;not null"`
}

type AdminUserRepoImpl struct {
	db *DB
}

func (AdminUser) TableName() string {
	return "admin_users"
}

func NewAdminUserRepoImpl(db *DB) *AdminUserRepoImpl {
	return &AdminUserRepoImpl{
		db: db,
	}
}

func (impl *AdminUserRepoImpl) EmailExist(email string, id uint) bool {
	query := impl.db.GetClient().Model(&AdminUser{}).Where("email = ?", email)
	if id > 0 {
		query = query.Where("id != ?", id)
	}
	var count int64
	err := query.Count(&count).Error
	if err != nil {
		return false // 出错时返回 false，让上层处理错误
	}
	return count > 0
}

func (impl *AdminUserRepoImpl) CreateUser(user *entity.User) (*entity.User, error) {
	adminUser := &AdminUser{
		Email:    user.Email,
		Password: user.Password,
		Status:   user.Status,
		Name:     user.Email, // 默认使用邮箱作为名称
	}
	if err := impl.db.GetClient().Create(adminUser).Error; err != nil {
		return nil, err
	}
	user.Id = uint(adminUser.Id)
	return user, nil
}
