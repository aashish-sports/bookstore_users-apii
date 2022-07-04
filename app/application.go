package app

import (
	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

func StartApplication() {
	mapurls()
	router.Run(":8080")

}
