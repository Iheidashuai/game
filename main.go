package main

import (
	"game/model"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, model.NewTestCities())
	})

	r.Run(":8080")
}
