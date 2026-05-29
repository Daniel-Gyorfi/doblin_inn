package api

import (
	"github.com/gin-gonic/gin"
)

func Routes(server *gin.Engine) {
	ImgUpload(server)
}
