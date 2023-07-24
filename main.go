package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

const response1 = "5DD7D16FE42EBFF9F717B71559C05D13D9F3C89DDD28AF48902984053D"
const response2 = "1t21Pz6tHt29uIt2ayHR6M8LKovW5fiddG+m0ds="

func main() {
	r := gin.Default()

	r.GET("/control.php", func(c *gin.Context) {
		c.Data(http.StatusOK, "application/json; charset=utf-8", []byte(response1))
	})

	r.GET("/control2.php", func(c *gin.Context) {
		c.Data(http.StatusOK, "application/json; charset=utf-8", []byte(response2))
	})

	r.GET("/control3.php", func(c *gin.Context) {
		c.Data(http.StatusOK, "application/json; charset=utf-8", []byte(response1))
	})

	r.Run(":80")
}
