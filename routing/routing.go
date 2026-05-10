package routing

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gitlab.com/DG_Blaster/doblin_inn/models"
	view "gitlab.com/DG_Blaster/doblin_inn/templates"
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

func Routes(server *gin.Engine) {
	About(server)
	Reserve(server)
	Locate(server)
}
