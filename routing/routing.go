package routing

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	controller "gitlab.com/DG_Blaster/doblin_inn/controller"
	view "gitlab.com/DG_Blaster/doblin_inn/templates"
	"gorm.io/gorm"
)

func Reserve(server *gin.Engine) {
	server.GET("/reserve", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "", view.Layout(view.Reserve()))
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
				} else {
					// ctx.String(http.StatusInternalServerError, "Database error")
					ctx.HTML(http.StatusInternalServerError, "", view.Layout(view.ErrorInternalServer()))
					return
				}
			}

			// Render the templ page passing the retrieved model
			ctx.HTML(http.StatusOK, "", view.Layout(view.Location_Index(location)))
		})
}

func Routes(server *gin.Engine) {
	view.About(server)
	Reserve(server)
	view.Locate(server)
	Location_Index(server)
	view.Login(server)
}
