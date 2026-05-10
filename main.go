package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	models "gitlab.com/DG_Blaster/doblin_inn/models"
	"gitlab.com/DG_Blaster/doblin_inn/routing"
	gintemplrenderer "gitlab.com/DG_Blaster/doblin_inn/templ_renderer"
	view "gitlab.com/DG_Blaster/doblin_inn/templates"
)

func main() {
	// 1. Connect & seed DB first
	models.ConnectDatabase()
	models.InitializeDatabase()

	//initialize server and html
	server := gin.New()

	server.SetTrustedProxies(nil)
	server.Static("/static", "./static")
	server.Static("/assets", "./assets")
	// server.
	server.HTMLRender = gintemplrenderer.Default

	//index
	server.GET("/", showHome)
	routing.Routes(server)

	server.Run() // specifying port 8080 here causes an error for some reason

}

// index page, other routes are in the routing folder
func showHome(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "", view.Layout(view.Index()))
}
