package models

import (
	"time"
)

type Location struct {
	ID         uint      `gorm:"primaryKey" json:"id"`
	Name       string    `gorm:"not null"   json:"name"`
	Img        string    `json:"img"`
	Descript   string    `json:"descript"`
	CreatedAt  time.Time `json:"createdAt"`
	ModifiedAt time.Time `json:"modifiedAt"`
}
