package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Go app...")

	//initialize server and html
	server := gin.New()
	server.LoadHTMLGlob("./templates/*.html")

	//index
	server.GET("/", showHome)

	server.Run("8080")
}

// index page
func showHome(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "index.html", nil)
}
