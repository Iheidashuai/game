package main

import (
	"game/db"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	d, err := db.NewDB()
	if err != nil {
		panic(err)
	}
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, &struct {
			Id   int64  `json:"id"`
			Name string `json:"name"`
		}{
			Id:   1,
			Name: "100",
		})
	})

	r.GET("/user", func(c *gin.Context) {
		user, err := d.UserByUid(c, 1)
		if err != nil {
			c.JSON(http.StatusInternalServerError, nil)
			return
		}
		c.JSON(http.StatusOK, user)
	})

	r.Run(":8080")
}
