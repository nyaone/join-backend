package admin

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"join-nyaone/global"
	"join-nyaone/models"
	"net/http"
)

func InviteeList(ctx *gin.Context) {
	userId, exist := ctx.Get("UID")
	if !exist {
		global.Logger.Errorf("No UID provided by middleware")
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "Middleware not working properly",
		})
		return
	}

	var invitees []models.User
	responseInvitees := []InviteesResponse{} // So it won't be null
	// Prepare base query
	query := global.DB

	// Check if code is specified
	codeUUIDStr := ctx.Param("code")
	if codeUUIDStr != "" {
		// Code is specified
		codeUUID, err := uuid.Parse(codeUUIDStr)
		if err != nil {
			global.Logger.Errorf("Failed to parse provided code (%s) with error: %v", codeUUIDStr, err)
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": "Invalid code format",
			})
			return
		}

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

		// Add additional conditions
		query = query.Where("invited_by_code = ?", targetCode.Code.Code)
	} else {
		// Just user (list all)
		query = query.Where("invited_by_user_id = ?", userId.(uint))
	}

	// Find results
	query.Find(&invitees)

	for _, invitee := range invitees {
		responseInvitees = append(responseInvitees, InviteesResponse{
			RegisteredAt:  invitee.CreatedAt,
			Username:      invitee.Username,
			InvitedByCode: invitee.InvitedByCode.String(),
		})
	}

	ctx.JSON(http.StatusOK, responseInvitees)

}
