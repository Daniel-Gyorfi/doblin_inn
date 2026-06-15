package controller

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"gitlab.com/DG_Blaster/doblin_inn/models"
	"gorm.io/gorm"
)

type LocationImageController struct {
	DB *gorm.DB
}

func NewLocationImageController(db *gorm.DB) *LocationImageController {
	return &LocationImageController{DB: db}
}

// CreateLocationImageInput defines the request body for creating an image
type CreateLocationImageInput struct {
	LocID    int    `json:"locId" binding:"required"`
	Img      string `json:"img"`
	Descript string `json:"descript"`
}

// UpdateLocationImageInput defines the request body for updating an image
type UpdateLocationImageInput struct {
	LocID    int    `json:"locId"`
	Img      string `json:"img"`
	Descript string `json:"descript"`
}

// Create handles POST /location-images
func (ctrl *LocationImageController) Create(ctx *gin.Context) {
	var input CreateLocationImageInput
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newImage := models.LocationImage{
		LocID:      input.LocID,
		Img:        input.Img,
		Descript:   input.Descript,
		CreatedAt:  time.Now(),
		ModifiedAt: time.Now(),
	}

	if err := ctrl.DB.Create(&newImage).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, newImage)
}

// GetAll handles GET /location-images
func (ctrl *LocationImageController) GetAll(ctx *gin.Context) {
	var images []models.LocationImage
	if err := ctrl.DB.Find(&images).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, images)
}

// GetByID handles GET /location-images/:id
func (ctrl *LocationImageController) GetByID(ctx *gin.Context) {
	id := ctx.Param("id")
	var image models.LocationImage

	if err := ctrl.DB.First(&image, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "Location image not found"})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, image)
}

// GetByLocationID handles GET /locations/:locId/images
func (ctrl *LocationImageController) GetByLocationID(ctx *gin.Context) {
	locID := ctx.Param("locId")
	var images []models.LocationImage

	if err := ctrl.DB.Where("loc_id = ?", locID).Find(&images).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, images)
}

// Update handles PUT /location-images/:id
func (ctrl *LocationImageController) Update(ctx *gin.Context) {
	id := ctx.Param("id")
	var image models.LocationImage

	// Check if the record exists
	if err := ctrl.DB.First(&image, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "Location image not found"})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var input UpdateLocationImageInput
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Update fields if provided
	if input.LocID != 0 {
		image.LocID = input.LocID
	}
	if input.Img != "" {
		image.Img = input.Img
	}
	if input.Descript != "" {
		image.Descript = input.Descript
	}
	image.ModifiedAt = time.Now()

	if err := ctrl.DB.Save(&image).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, image)
}

// Delete handles DELETE /location-images/:id
func (ctrl *LocationImageController) Delete(ctx *gin.Context) {
	id := ctx.Param("id")
	var image models.LocationImage

	if err := ctrl.DB.First(&image, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "Location image not found"})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if err := ctrl.DB.Delete(&image).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Location image deleted successfully"})
}
