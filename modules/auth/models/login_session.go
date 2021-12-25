package models

import "time"

type LoginSession struct {
	CreatedAt  time.Time
	UpdatedAt  time.Time
	LastUsed   time.Time
	K          string `gorm:"type:varchar(64);column:k;unique" json:"key"`
	V          string `gorm:"type:bytea;column:v" json:"value"`
	E          int64  `gorm:"column:e;default:0;" json:"expire_ttl"`
	ID         uint   `gorm:"primaryKey" json:"id"`
	UserID     uint   `gorm:"user_id" json:"user_id"`
	RememberMe bool   `json:"remember_me" gorm:"remember_me" form:"remember_me"`
}
