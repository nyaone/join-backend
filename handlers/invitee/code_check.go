package invitee

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"join-nyaone/global"
	"join-nyaone/models"
	"join-nyaone/utils"
	"net/http"
)

func CodeCheck(ctx *gin.Context) {

	codeUUIDStr := ctx.Param("code")
	valid, errCode, _, invalidReason, err := CheckInviteCodeStatus(codeUUIDStr)
	if err != nil {
		ctx.JSON(errCode, gin.H{
			"error": err.Error(),
		})
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"valid":  valid,
			"reason": invalidReason,
		})
	}

}

func CheckInviteCodeStatus(codeUUIDStr string) (bool, int, *models.Code, string, error) {
	// 1. Check format
	codeUUID, err := uuid.Parse(codeUUIDStr)
	if err != nil {
		global.Logger.Errorf("Failed to parse provided code (%s) with error: %v", codeUUIDStr, err)
		return false, http.StatusBadRequest, nil, "Invalid code format", fmt.Errorf("invalid code format")
	}

	// 2. Check exist
	var targetCode models.Code
	err = global.DB.First(&targetCode, "code = ?", codeUUID).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return false, http.StatusNotFound, nil, "No such code", nil
	} else if err != nil {
		global.Logger.Errorf("Failed to find code into database with error: %v", err)
		return false, http.StatusInternalServerError, nil, "Failed to find code", fmt.Errorf("failed to find code")
	}

	// 3. Check valid
	isValid, _, invalidReason := utils.CheckInviteCodeValid(&targetCode)

	return isValid, 0, &targetCode, invalidReason, nil
}
