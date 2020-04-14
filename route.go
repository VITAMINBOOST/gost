package main

import (
	"gost/gostaudio"
	"net/http"

	"github.com/gin-gonic/gin"
)

func InitRouter() {
	v1 := r.Group("v1")

	r.GET("/uploadpage", func(c *gin.Context) {
		title := "upload single file"
		c.HTML(http.StatusOK, "upload.html", gin.H{
			"page": title,
		})
	})

	r.GET("/multiuploadpage", func(c *gin.Context) {
		title := "multiple upload single file"
		c.HTML(http.StatusOK, "multiupload.html", gin.H{
			"page": title,
		})
	})

	audios := v1.Group("audio")
	{

		audios.POST("/upload", gostaudio.SingleUploadFromClient)
		audios.POST("/multiupload", gostaudio.MultiUploadFromClient)

	}
}
