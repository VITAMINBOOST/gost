package gostrouter

import (
	"github.com/gin-gonic/gin"
)

func InitRouter(r *gin.Engine) {
	v1 := r.Group("v1")

	audios := v1.Group("audio")
	{
		audios.GET("/upload", gostaudio.)
	}

}
