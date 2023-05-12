package app

import (
	//"log"

	//"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"net/http"

)

var (
	router *gin.Engine
)

func init() {
	router = gin.Default()
}

func StartApp() {
	
	router.Use(func(c *gin.Context) {
        c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
        c.Writer.Header().Set("Access-Control-Allow-Methods", "OPTIONS, GET, POST, PUT, DELETE")
        c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type")
        if c.Request.Method == http.MethodOptions {
            c.AbortWithStatus(http.StatusNoContent)
            return
        }
        c.Next()
    })
	mapUrls()

	if err := router.Run(":25025"); err != nil {
		panic(err)
	}
}