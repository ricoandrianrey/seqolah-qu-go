package models

import (
	"seqolah-qu/helpers"
	"seqolah-qu/types"
	"time"

	"gorm.io/gorm"
)

type User struct {
	Id int `json:"id" gorm:"primaryKey;autoIncrement"`
	// SchoolId  int            `json:"school_id"`
	Email     string         `json:"email" gorm:"type:varchar;uniqueIndex"`
	Phone     string         `json:"phone" gorm:"type:varchar(20)"`
	Password  string         `json:"password" gorm:"type:varchar"`
	Token     string         `json:"token" gorm:"type:varchar"`
	OtpConfig types.JSONB    `json:"otp_config" gorm:"type:jsonb"`
	Status    string         `json:"status" gorm:"type:varchar"`
	Role      string         `json:"role" gorm:"type:varchar(25)"` // student | school | parent
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
	// School    School         `json:"school"`
}

func (u *User) BeforeCreate(tx *gorm.DB) error {
	u.Password = helpers.GenerateHashPassword(u.Password)
	return nil
}
