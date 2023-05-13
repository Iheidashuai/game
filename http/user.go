package http

import (
	"strconv"

	"game/service"
	"github.com/gin-gonic/gin"
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

func StrongerEquipment(ctx *gin.Context) {
	userId := ctx.Param("user_id")
	// 转换成 int64
	userIdInt64, err := strconv.ParseInt(userId, 10, 64)
	if err != nil {
		return
	}

	weaponId := ctx.Param("weapon_id")
	// 转换成 int64
	weaponIdInt64, err := strconv.ParseInt(weaponId, 10, 64)
	if err != nil {
		return
	}

	ctx.JSON(200, service.StrongerEquipment(ctx, userIdInt64, weaponIdInt64).Error())
}
