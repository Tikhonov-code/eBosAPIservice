package utils

import (
	"log"

	"github.com/gin-gonic/gin"
)

func Respond(c *gin.Context, status int, body interface{}) {

	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Headers", "Content-Type, Authorization")
	c.Header("Access-Control-Allow-Methods","OPTIONS,GET,POST,PUT,PATCH,DELETE");


	if c.GetHeader("Accept") == "application/json" {
		log.Println("Respond XML")
		c.XML(status, body)
		return
	}
	log.Println("Respond json")
	c.JSON(status,body)

}
func RespondErr(c *gin.Context, err *ApplicationError) {
	
	log.Println("RespondErr",err.Message)

	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Headers", "Content-Type, Authorization")
    c.Header("Access-Control-Allow-Methods","OPTIONS,GET,POST,PUT,PATCH,DELETE");

	if c.GetHeader("Accept") == "application/json" {
		c.XML(err.StatusCode, err)
		return
	}
	c.JSON(err.StatusCode, err)

}

