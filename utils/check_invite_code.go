package utils

import (
	"context"
	"fmt"
	"join-nyaone/consts"
	"join-nyaone/global"
	"join-nyaone/models"
	"time"
)

func CheckInviteCodeValid(targetCode *models.Code) (bool, int64) {
	var registeredCountWithThisCode int64
	global.DB.Model(&models.User{}).Where("invited_by_code = ?", targetCode.Code.Code).Count(&registeredCountWithThisCode)

	//// 3.1. Check if disabled by admin
	if !targetCode.IsActivate {
		// Inactive
		return false, registeredCountWithThisCode
	}
	//// 3.2. Check time period
	nowTime := time.Now()
	if targetCode.RegisterTimeStart.After(nowTime) || (targetCode.IsRegisterTimeEndValid && targetCode.RegisterTimeEnd.Before(nowTime)) {
		// Exceeds acceptable time
		return false, registeredCountWithThisCode
	}
	//// 3.3. Check not exceeds max register limit
	if targetCode.RegisterCountLimit > 0 {

		if registeredCountWithThisCode >= targetCode.RegisterCountLimit {
			// Exceeds maximum register limit
			return false, registeredCountWithThisCode
		}
	}

	//// 3.4. Check if enables cool-down, not in cd
	if targetCode.RegisterCoolDown > 0 {
		// Check in redis
		redisKey := fmt.Sprintf(consts.REDIS_KEY_REGISTER_CODE_CD, targetCode.Code.Code.String())
		if exist, err := global.Redis.Exists(context.Background(), redisKey).Result(); err != nil {
			// Failed to check
			global.Logger.Errorf("Failed to check code cd in redis")
			// Maybe just ignore?
			return true, registeredCountWithThisCode
		} else if exist > 0 {
			// Nope, it's in cd
			return false, registeredCountWithThisCode
		}
	}

	return true, registeredCountWithThisCode
}
