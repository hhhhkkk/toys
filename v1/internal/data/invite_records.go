package data

import (
	"time"

	"github.com/hhhhkkk/mini-blog/v1/internal/biz/api/entity"
)

type UserInviteRecord struct {
	Id         int       `gorm:"primaryKey"`
	InviteUid  int       `gorm:"type:int(11);not null"`
	InvitedUId int       `gorm:"type:int(11);not null"`
	CreatedAt  time.Time `gorm:"type:datetime;not null"`
	UpdatedAt  time.Time `gorm:"type:datetime;not null"`
}

func (u *UserInviteRecord) TableName() string {
	return "user_invite_records"
}

type InviteRecordRepoImpl struct {
	db *DB
}

func NewInviteRecordRepoImpl(db *DB) *InviteRecordRepoImpl {
	return &InviteRecordRepoImpl{
		db: db,
	}
}

func (impl *InviteRecordRepoImpl) CreateInvite(en *entity.InviteEntity) (*entity.InviteEntity, error) {
	userInviteRecord := &UserInviteRecord{
		InviteUid:  en.InviteUid,
		InvitedUId: en.InvitedUid,
	}
	if err := impl.db.GetClient().Create(userInviteRecord).Error; err != nil {
		return nil, err
	}
	en.Id = userInviteRecord.Id
	en.CreatedAt = userInviteRecord.CreatedAt
	en.UpdatedAt = userInviteRecord.UpdatedAt
	return en, nil
}

func (impl *InviteRecordRepoImpl) GetByInvitedUid(uid int) *entity.InviteEntity {
	var userInviteRecord UserInviteRecord
	tx := impl.db.GetClient().Model(&UserInviteRecord{}).Where("invited_uid = ?", uid).First(&userInviteRecord)
	if tx.Error != nil {
		return nil
	}
	return &entity.InviteEntity{
		Id:         userInviteRecord.Id,
		InviteUid:  userInviteRecord.InviteUid,
		InvitedUid: userInviteRecord.InvitedUId,
		CreatedAt:  userInviteRecord.CreatedAt,
		UpdatedAt:  userInviteRecord.UpdatedAt,
	}
}

func (impl *InviteRecordRepoImpl) GetByInviteUid(uid int) []*entity.InviteEntity {
	var userInviteRecords []UserInviteRecord
	tx := impl.db.GetClient().Model(&UserInviteRecord{}).Where("invite_uid = ?", uid).Find(&userInviteRecords)
	if tx.Error != nil {
		return nil
	}
	var entities []*entity.InviteEntity
	for _, userInviteRecord := range userInviteRecords {
		entities = append(entities, &entity.InviteEntity{
			Id:         userInviteRecord.Id,
			InviteUid:  userInviteRecord.InviteUid,
			InvitedUid: userInviteRecord.InvitedUId,
			CreatedAt:  userInviteRecord.CreatedAt,
			UpdatedAt:  userInviteRecord.UpdatedAt,
		})
	}
	return entities
}

func (impl *InviteRecordRepoImpl) GetInviteTotal(uid int) int {
	var count int64
	tx := impl.db.GetClient().Model(&UserInviteRecord{}).Where("invite_uid = ?", uid).Count(&count)
	if tx.Error != nil {
		return 0
	}
	return int(count)
}
