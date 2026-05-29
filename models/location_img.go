package models

import (
	"time"
)

type LocationImage struct {
	ID         int       `gorm:"primaryKey" json:"id"`
	LocID      int       `gorm: "foreignKey" json:"id"`
	Img        string    `json:"img"`
	Descript   string    `json:"descript"`
	CreatedAt  time.Time `json:"createdAt"`
	ModifiedAt time.Time `json:"modifiedAt"`
}
