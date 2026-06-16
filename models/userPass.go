package models

import "time"

type UserPass struct {
	ID         int       `gorm:"primaryKey" json:"id"`
	UserID     int       `gorm:"foreignKey" json:"userId"`
	CreatedAt  time.Time `gorm:"not null"	json:"createdAt"`
	ModifiedAt time.Time `json:"modifiedAt"`
	passHash   string    `gorm:"not null" json:passHass`
}
