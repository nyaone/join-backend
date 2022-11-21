package admin

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"join-nyaone/global"
	"join-nyaone/models"
	"join-nyaone/types"
	"join-nyaone/utils"
	"net/http"
)

func CodeEdit(ctx *gin.Context) {
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

	// Parse code string to uuid
	codeUUIDStr := ctx.Param("code")
	codeUUID, err := uuid.Parse(codeUUIDStr)
	if err != nil {
		global.Logger.Errorf("Failed to parse provided code (%s) with error: %v", codeUUIDStr, err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid code format",
		})
		return
	}

	// Find code with provided code-code
	var targetCode models.Code
	err = global.DB.First(&targetCode, "code = ? AND created_by_user_id = ?", codeUUID, userId).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": "No such code",
		})
		return
	} else if err != nil {
		global.Logger.Errorf("Failed to find code into database with error: %v", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to find code",
		})
		return
	}

	// Update props
	targetCode.CodeProps = props

	err = global.DB.Save(&targetCode).Error
	if err != nil {
		global.Logger.Errorf("Failed to save code into database with error: %v", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to save code",
		})
		return
	}

	// All success
	isValid, inviteCount := utils.CheckInviteCodeValid(&targetCode)
	ctx.JSON(http.StatusOK, CodeResponse{
		Code:        targetCode.Code.Code.String(),
		CodeProps:   targetCode.CodeProps,
		InviteCount: inviteCount, // Ignored here
		IsValid:     isValid,
	})
}
