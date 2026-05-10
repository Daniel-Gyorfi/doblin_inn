package controller

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"gitlab.com/DG_Blaster/doblin_inn/models"
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

func FindLocations(c *gin.Context) {
	var locations []models.Location
	models.DB.Find(&locations)

	c.JSON(http.StatusOK, gin.H{"data": locations})
}

func FindLocation(c *gin.Context) {
	var location models.Location
	if err := models.DB.Where("id = ?", c.Param("id")).First(&location).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": location})
}

func CreateLocation(c *gin.Context) {
	var input CreateLocationInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
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

	c.JSON(http.StatusOK, gin.H{"data": location})
}

func UpdateLocation(c *gin.Context) {
	var location models.Location
	if err := models.DB.Where("id = ?", c.Param("id")).First(&location).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	var input UpdateLocationInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
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

	c.JSON(http.StatusOK, gin.H{"data": location})
}

func DeleteLocation(c *gin.Context) {
	var location models.Location
	if err := models.DB.Where("id = ?", c.Param("id")).First(&location).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	models.DB.Delete(&location)

	c.JSON(http.StatusOK, gin.H{"data": true})
}
