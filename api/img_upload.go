package api

import (
	"fmt"
	"log"
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

func ImgUpload(server *gin.Engine) {
	server.POST("/upload/loc_img", func(ctx *gin.Context) {
		// single file
		file, err := ctx.FormFile("file")
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		log.Println(file.Filename)

		// Upload the file to specific dst.
		dst := filepath.Join("./assets/uploads/", filepath.Base(file.Filename))
		ctx.SaveUploadedFile(file, dst)

		ctx.String(http.StatusOK, fmt.Sprintf("'%s' uploaded!", file.Filename))
	})
}
