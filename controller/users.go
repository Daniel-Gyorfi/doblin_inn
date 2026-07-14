package controller

import (
	"errors"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"gitlab.com/DG_Blaster/doblin_inn/models"
	"gorm.io/gorm"
)

type CreateUserInput struct {
	Username   string    `json:"username" binding:"required"`
	CreatedAt  time.Time `json:"createdAt"`
	ModifiedAt time.Time `json:"modifiedAt"`
}

type UpdateUserInput struct {
	Username   string    `json:"username"`
	CreatedAt  time.Time `json:"createdAt"`
	ModifiedAt time.Time `json:"modifiedAt"`
}

func FindUsers(ctx *gin.Context) {
	var users []models.User
	models.DB.Find(&users)

	ctx.JSON(http.StatusOK, gin.H{"data": users})
}

func GetUserByID(id interface{}) (models.User, error) {
	var user models.User
	err := models.DB.Where("id = ?", id).First(&user).Error
	return user, err
}

func FindUser(ctx *gin.Context) {
	user, err := GetUserByID(ctx.Param("id"))

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "Record not found!"})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": user})
}

func DeleteUser(ctx *gin.Context) {
	user, err := GetUserByID(ctx.Param("id"))

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	models.DB.Delete(&user)

	ctx.JSON(http.StatusOK, gin.H{"data": true})
}
