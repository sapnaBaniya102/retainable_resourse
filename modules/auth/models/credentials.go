package models

type Credential struct {
	Password []byte `gorm:"type:bytea;column:password" json:"value"`
	BaseModel
	UserID uint  `gorm:"user_id" json:"user_id"`
	E      int64 `gorm:"column:e;default:0;" json:"expire_ttl"`
}
