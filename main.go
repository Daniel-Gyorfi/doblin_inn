package main

import (
	"net/http"

	"github.com/Daniel-Gyorfi/doblin_inn/routing"
	gintemplrenderer "github.com/Daniel-Gyorfi/doblin_inn/templ_renderer"
	view "github.com/Daniel-Gyorfi/doblin_inn/templates"
	"github.com/gin-gonic/gin"
)

func main() {

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
