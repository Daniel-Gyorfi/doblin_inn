package models

import (
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var DB *gorm.DB

func ConnectDatabase() {
	database, err := gorm.Open("sqlite3", "sqlite.db")

	if err != nil {
		panic("Failed to connect to database!")
	}

	database.AutoMigrate(&Location{})

	DB = database
}

func InitializeDatabase() {
	// Only seed if the table is empty
	var count int64
	DB.Model(&Location{}).Count(&count)
	if count > 0 {
		return
	}

	now := time.Now()
	seed := []Location{
		{
			Name:       "Savior's Row",
			Img:        "saviors_row.png",
			Descript:   "A proud district now riddled with secrets and glowing streets.",
			CreatedAt:  now,
			ModifiedAt: now,
		},
		{
			Name:       "Dobarum Keep",
			Img:        "landing_hotel.jpg",
			Descript:   "Within a gleaming stone fortress overlooking the Skysdale historical district",
			CreatedAt:  now,
			ModifiedAt: now,
		},
		{
			Name:       "Old Samnium",
			Img:        "old_samnium.png",
			Descript:   "Sun-bleached outpost where merchants still barter among fallen houses.",
			CreatedAt:  now,
			ModifiedAt: now,
		},
	}

	for _, loc := range seed {
		DB.Create(&loc)
	}
}
