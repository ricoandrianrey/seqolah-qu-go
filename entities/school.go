package entities

import (
	"seqolah-qu/types"
	"time"

	"gorm.io/gorm"
)

type School struct {
	ID            uint           `gorm:"primaryKey;autoIncrement" json:"id"`
	Address       string         `gorm:"type:text" json:"address"`
	AddressDetail types.JSONB    `gorm:"type:jsonb" json:"address_detail"`
	Logo          string         `gorm:"type:varchar" json:"logo"`
	Email         string         `gorm:"type:varchar" json:"email"`
	Accreditation string         `gorm:"type:varchar" json:"accreditation"`
	Phone         string         `gorm:"type:varchar(15)" json:"phone"`
	Curriculum    string         `gorm:"type:varchar" json:"curriculum"`
	CreatedAt     time.Time      `json:"created_at"`
	UpdatedAt     time.Time      `json:"updated_at"`
	DeletedAt     gorm.DeletedAt `json:"deleted_at"`
}
