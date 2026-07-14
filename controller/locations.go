package controller

import (
	"errors"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"gitlab.com/DG_Blaster/doblin_inn/models"
	"gorm.io/gorm"
)

type CreateLocationInput struct {
	Name     string `json:"name" binding:"required"`
	Img      string `json:"img"`
	Descript string `json:"descript"`
}

type UpdateLocationInput struct {
	Name     string `json:"name"`
	Img      string `json:"img"`
	Descript string `json:"descript"`
}

func FindLocations(ctx *gin.Context) {
	var locations []models.Location
	models.DB.Find(&locations)

	ctx.JSON(http.StatusOK, gin.H{"data": locations})
}

func GetLocationByID(id interface{}) (models.Location, error) {
	var location models.Location
	err := models.DB.Where("id = ?", id).First(&location).Error
	return location, err
}

func FindLocation(ctx *gin.Context) {
	location, err := GetLocationByID(ctx.Param("id"))
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "Record not found!"})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": location})
}

func CreateLocation(ctx *gin.Context) {
	var input CreateLocationInput
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	now := time.Now()
	location := models.Location{
		Name:       input.Name,
		Img:        input.Img,
		Descript:   input.Descript,
		CreatedAt:  now,
		ModifiedAt: now,
	}
	models.DB.Create(&location)

	ctx.JSON(http.StatusOK, gin.H{"data": location})
}

func UpdateLocation(ctx *gin.Context) {
	var location models.Location
	if err := models.DB.Where("id = ?", ctx.Param("id")).First(&location).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	var input UpdateLocationInput
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updates := map[string]interface{}{
		"ModifiedAt": time.Now(),
	}
	if input.Name != "" {
		updates["Name"] = input.Name
	}
	if input.Img != "" {
		updates["Img"] = input.Img
	}
	if input.Descript != "" {
		updates["Descript"] = input.Descript
	}

	models.DB.Model(&location).Updates(updates)

	ctx.JSON(http.StatusOK, gin.H{"data": location})
}

func DeleteLocation(ctx *gin.Context) {

	location, err := GetLocationByID(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	models.DB.Delete(&location)

	ctx.JSON(http.StatusOK, gin.H{"data": true})
}
