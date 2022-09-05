package entities

import (
	"time"

	"gorm.io/gorm"
)

type UserParent struct {
	Id        int            `json:"id" gorm:"primaryKey;autoIncrement"`
	UserId    int            `json:"user_id"`
	Name      string         `json:"name" gorm:"type:varchar"`
	Gender    string         `json:"gender" gorm:"type:varchar"`
	Photo     string         `json:"photo" gorm:"type:varchar"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
	User      User           `json:"user"`
	// UserStudents []UserStudent  `json:"user_students" gorm:"foreignKey:ParentId;references:Id"`
}
