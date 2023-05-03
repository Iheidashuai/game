package http

import (
	"game/service"
	"github.com/gin-gonic/gin"
	"strconv"
)

func GetUser(ctx *gin.Context) {
	userId := ctx.Param("user_id")
	// 转换成 int64
	userIdInt64, err := strconv.ParseInt(userId, 10, 64)
	if err != nil {
		return
	}

	user, err := service.GetUser(ctx, userIdInt64)
	if err != nil {
		return
	}
	ctx.JSON(200, user)
}
