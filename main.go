package main

import (
	"game/http"
	"game/model"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, model.NewTestCities())
	})

	r.GET("/user/:user_id", http.GetUser)

	r.Run(":8080")
}
