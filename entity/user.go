package entity

import "time"

type User struct {
	ID                 int       `gorm:"primaryKey" json:"id"`
	Username           string    `json:"username" gorm:"type:varchar(255);not null"`
	Email              string    `json:"email" gorm:"type:varchar(255);not null"`
	Password           string    `json:"password" gorm:"type:varchar(255);not null"`
	StatusVerification string    `json:"status_verification" gorm:"type:varchar(255);null"`
	CreatedAt          time.Time `json:"created_at"`
	UpdatedAt          time.Time `json:"updated_at"`
}
