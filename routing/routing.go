package routing

import (
	"net/http"

	view "github.com/Daniel-Gyorfi/doblin_inn/templates"
	"github.com/gin-gonic/gin"
)

func About(server *gin.Engine) {
	server.GET("/about", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "", view.Layout(view.About()))
	})
}
