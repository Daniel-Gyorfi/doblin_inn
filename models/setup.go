package models

import (
	"log"
	"os"
	"time"

	"gorm.io/driver/sqlite" // GORM v2 wrapper
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	// make sure the folder is present
	if err := os.MkdirAll("data", 0755); err != nil {
		log.Fatalf("cannot create data dir: %v", err)
	}

	database, err := gorm.Open(sqlite.Open("data/sqlite.db"), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
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
			Descript:   "A proud district now riddled with revanchists.",
			CreatedAt:  now,
			ModifiedAt: now,
		},
		{
			Name:       "Dobarum Keep",
			Img:        "landing_hotel.jpg",
			Descript:   "Within a gleaming citadel overlooking the Skysdale historical district",
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
