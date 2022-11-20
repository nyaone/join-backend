package invitee

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"join-nyaone/global"
	"join-nyaone/models"
	"net/http"
	"time"
)

func CodeCheck(ctx *gin.Context) {

	codeUUIDStr := ctx.Param("code")
	valid, errCode, _, err := CheckInviteCodeStatus(codeUUIDStr)
	if err != nil {
		ctx.JSON(errCode, gin.H{
			"error": err.Error(),
		})
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"valid": valid,
		})
	}

}

func CheckInviteCodeStatus(codeUUIDStr string) (bool, int, *models.Code, error) {
	// 1. Check format
	codeUUID, err := uuid.Parse(codeUUIDStr)
	if err != nil {
		global.Logger.Errorf("Failed to parse provided code (%s) with error: %v", codeUUIDStr, err)
		return false, http.StatusBadRequest, nil, fmt.Errorf("invalid code format")
	}

	// 2. Check exist
	var targetCode models.Code
	err = global.DB.First(&targetCode, "code = ?", codeUUID).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
	} else if err != nil {
		global.Logger.Errorf("Failed to find code into database with error: %v", err)
		return false, http.StatusInternalServerError, nil, fmt.Errorf("failed to find code")
	}

	// 3. Check valid
	isValid, _ := CheckInviteCodeValid(&targetCode)

	return isValid, 0, &targetCode, nil
}

func CheckInviteCodeValid(targetCode *models.Code) (bool, int64) {
	//// 3.1. Check if disabled by admin
	if !targetCode.IsActivate {
		// Inactive
		return false, 0
	}
	//// 3.2. Check time period
	nowTime := time.Now()
	if targetCode.RegisterTimeStart.After(nowTime) || (targetCode.IsRegisterTimeEndValid && targetCode.RegisterTimeEnd.Before(nowTime)) {
		// Exceeds acceptable time
		return false, 0
	}
	//// 3.3. Check not exceeds max register limit
	var registeredCountWithThisCode int64
	if targetCode.RegisterCountLimit > 0 {
		global.DB.Model(&models.User{}).Where("invited_by_code = ?", targetCode.Code.Code).Count(&registeredCountWithThisCode)

		if registeredCountWithThisCode >= targetCode.RegisterCountLimit {
			// Exceeds maximum register limit
			return false, registeredCountWithThisCode
		}
	}

	return true, registeredCountWithThisCode
}
