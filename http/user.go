package http

import (
	"game/db"
	"github.com/gin-gonic/gin"
	"strconv"
)

func GetUser(ctx *gin.Context) {
	db, err := db.NewDB()
	if err != nil {
		return
	}

	userId := ctx.Param("user_id")
	// 转换成 int64
	userIdInt64, err := strconv.ParseInt(userId, 10, 64)
	if err != nil {
		return
	}
	user, err := db.UserByUid(ctx, userIdInt64)
	if err != nil {
		return
	}

	ctx.JSON(200, user)
}
