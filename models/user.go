package models

import "time"

type User struct {
	ID        int       `gorm:"primaryKey" json:"id"`
	Username  string    `gorm:"not null"   json:"username"`
	CreatedAt time.Time `json:"createdAt"`
}
