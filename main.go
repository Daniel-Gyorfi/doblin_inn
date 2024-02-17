package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {

	//initialize server and html
	server := gin.New()
	server.LoadHTMLGlob("./templates/*.html")

	//index
	server.GET("/", showHome)

	server.Run() // specifying port 8080 here causes an error for some reason
}

// index page, other routes are in the routing folder
func showHome(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "index.html", nil)
}
