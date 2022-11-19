package admin

import (
	"github.com/gin-gonic/gin"
	"join-nyaone/global"
	"join-nyaone/models"
	"join-nyaone/types"
	"net/http"
)

func CodeCreate(ctx *gin.Context) {
	userId, exist := ctx.Get("UID")
	if !exist {
		global.Logger.Errorf("No UID provided by middleware")
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "Middleware not working properly",
		})
		return
	}

	var props types.CodeProps
	err := ctx.ShouldBindJSON(&props)
	if err != nil {
		global.Logger.Errorf("Failed to parse provided props with error: %v", err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to parse provided props",
		})
		return
	}

	targetCode := models.Code{
		Code: types.Code{
			CreatedByUserID: userId.(uint),
			CodeProps:       props,
		},
	}

	err = global.DB.Create(&targetCode).Error
	if err != nil {
		global.Logger.Errorf("Failed to save code into database with error: %v", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to save code",
		})
		return
	}

	// All success
	ctx.JSON(http.StatusOK, CodeResponse{
		ID:          targetCode.ID,
		Code:        targetCode.Code.Code.String(),
		CodeProps:   targetCode.CodeProps,
		InviteCount: 0,
	})
}
