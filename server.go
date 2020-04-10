package main

import (
	"github.com/VITAMINBOOST/gost/gostrouter"
	"github.com/gin-gonic/gin"
)

var r *gin.Engine

func main() {
	r = gin.Default()

	r.Use(gin.Recovery())

	gostrouter.InitRouter()
	r.Run(":9900")

}
