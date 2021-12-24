package modules

import (
	"time"

	"gorm.io/gorm"
)

// GenerateIOSString generates a time string equivalent to the time.now
func GenerateIOSString() string {
	return time.Now().UTC().Format("2006-01-02T15:04:05.999Z07:00")

}

// Base contains common column for all tables
type Base struct {
	ID        uint   `gorm:"primaryKey" json:"id"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

//BeforeCreate will include Base struct after every insert
func (base *Base) BeforeCreate(tx *gorm.DB) error {

	//uuid.new() creates a new random UUIDor panics
	t := GenerateIOSString()
	base.CreatedAt, base.UpdatedAt = t, t
	return nil
}

//AfterUpdate will update the Base struct after every update
func (base *Base) AfterUpdate(tx *gorm.DB) error {

	//update timestamp
	base.UpdatedAt = GenerateIOSString()
	return nil
}
