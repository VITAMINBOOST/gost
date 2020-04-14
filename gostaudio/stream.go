package gostaudio

import "github.com/gin-gonic/gin"

func StreamAudio(c *gin.Context) {
	filename := c.PostForm("filename")
	
}