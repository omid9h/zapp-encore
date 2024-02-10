package userprofile

import "time"

type UserProfile struct {
	ID        string    `gorm:"type:uuid;primaryKey;default:uuid_generate_v4();not null"`
	UserID    string    `gorm:"type:uuid;not null"`
	Email     string    `gorm:"type:varchar(100);unique;unique_index;not null"`
	CreatedAt time.Time `gorm:"autoCreateTime:milli"`
	UpdatedAt time.Time `gorm:"autoUpdateTime:milli"`
	IsDeleted bool      `gorm:"default:false"`
}
