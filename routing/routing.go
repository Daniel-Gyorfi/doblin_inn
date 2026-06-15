package routing

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	controller "gitlab.com/DG_Blaster/doblin_inn/controller"
	"gitlab.com/DG_Blaster/doblin_inn/models"
	view "gitlab.com/DG_Blaster/doblin_inn/templates"
	"gorm.io/gorm"
)

func About(server *gin.Engine) {
	server.GET("/about", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "", view.Layout(view.About()))
	})
}

func Reserve(server *gin.Engine) {
	server.GET("/reserve", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "", view.Layout(view.Reserve()))
	})
}

func Locate(server *gin.Engine) {
	server.GET("/locations",
		func(ctx *gin.Context) {
			var locations []models.Location
			models.DB.Find(&locations)
			ctx.HTML(http.StatusOK, "", view.Layout(view.Locate(locations)))
		})
}

func Location_Index(server *gin.Engine) {
	server.GET("/locations/:id",
		func(ctx *gin.Context) {
			id := ctx.Param("id")

			// Use the controller helper logic to pull the model data
			location, err := controller.GetLocationByID(id)
			if err != nil {
				if errors.Is(err, gorm.ErrRecordNotFound) {
					// ctx.String(http.StatusNotFound, "Location not found")
					ctx.HTML(http.StatusNotFound, "", view.Layout(view.ErrorNotFound()))
					return
				}
				// ctx.String(http.StatusInternalServerError, "Database error")
				ctx.HTML(http.StatusInternalServerError, "", view.Layout(view.ErrorInternalServer()))
				return
			}

			// Render the templ page passing the retrieved model
			ctx.HTML(http.StatusOK, "", view.Layout(view.Location_Index(location)))
		})
}

func Login(server *gin.Engine) {
	server.GET("/log", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "", view.Layout(view.Login()))
	})
}

func Routes(server *gin.Engine) {
	About(server)
	Reserve(server)
	Locate(server)
	Location_Index(server)
	Login(server)
}
