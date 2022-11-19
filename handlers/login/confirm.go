package login

import (
	"context"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"join-nyaone/consts"
	"join-nyaone/global"
	"join-nyaone/models"
	"join-nyaone/utils"
	"net/http"
	"strings"
)

func Confirm(ctx *gin.Context) {
	token := ctx.Param("token")
	tokenSplits := strings.Split(token, "-")
	if len(tokenSplits) != 2 {
		// Invalid token format
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid token format.",
		})
		return
	}

	username := tokenSplits[0]
	secretFromRequest := tokenSplits[1]

	// Check if in redis
	loginSessionKey := fmt.Sprintf(consts.REDIS_KEY_LOGIN_REQUEST, username)
	if exist, err := global.Redis.Exists(context.Background(), loginSessionKey).Result(); err != nil {
		global.Logger.Errorf("Failed to check session key (%s) from redis with error: %v", loginSessionKey, err)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": fmt.Sprintf("Failed to check session from redis: %s", err.Error()),
		})
		return
	} else if exist == 0 {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"error": "Session doesn't exist.",
		})
		return
	}

	secretInServer, err := global.Redis.Get(context.Background(), loginSessionKey).Result()
	if err != nil {
		global.Logger.Errorf("Failed to get session key (%s) from redis with error: %v", loginSessionKey, err)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": fmt.Sprintf("Failed to get session from redis: %s", err.Error()),
		})
		return
	}

	if secretInServer != secretFromRequest {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"error": "Session secret incorrect.",
		})
		return
	}

	// Else: Valid
	global.Redis.Del(context.Background(), loginSessionKey)

	adminSession := utils.RandString(32)
	adminSessionKey := fmt.Sprintf(consts.REDIS_KEY_VALID_SESSION, adminSession)

	// Check user
	var user models.User
	err = global.DB.First(&user, "username = ?", username).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// Create record
			user.Username = username

			// Insert into database
			err = global.DB.Create(&user).Error
			if err != nil {
				global.Logger.Errorf("Failed to insert user (%s) into database with error: %v", username, err)
				ctx.JSON(http.StatusInternalServerError, gin.H{
					"error": "Failed to insert user into database.",
				})
				return
			}
		} else {
			global.Logger.Errorf("Failed to get user (%s) data from database with error: %v", username, err)
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error": "Failed to get user data from database.",
			})
			return
		}
	}

	// Save session into redis
	err = global.Redis.Set(context.Background(), adminSessionKey, user.ID, consts.TIME_LOGIN_SESSION_VALID).Err()
	if err != nil {
		global.Logger.Errorf("Failed to set session (%s) into redis with error: %v", adminSessionKey, err)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": fmt.Sprintf("Failed to set session into redis: %s", err.Error()),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"ok":       true,
		"session":  adminSession,
		"username": username,
	})

}
