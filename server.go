package main

import "github.com/gin-gonic/gin"

var r *gin.Engine

func main() {
	r = gin.Default()

	r.Use(gin.Recovery())
	InitRouter()

	r.MaxMultipartMemory = 8 << 20 // 8 MiB

	r.Static("/files", "./mediafiles/")
	r.LoadHTMLGlob("./templates/*")

	r.Run(":9000")
}
