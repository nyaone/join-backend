package login

import (
	"context"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"join-nyaone/consts"
	"join-nyaone/global"
	"join-nyaone/misskey"
	"join-nyaone/models"
	"join-nyaone/utils"
	"net/http"
	"time"
)

func Confirm(ctx *gin.Context) {
	token := ctx.Param("token")

	// Check if in redis
	loginSessionKey := fmt.Sprintf(consts.REDIS_KEY_LOGIN_REQUEST, token)
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

	// Retrieve saved application secret
	appSecret, err := global.Redis.Get(context.Background(), loginSessionKey).Result()
	if err != nil {
		global.Logger.Errorf("Failed to get session key (%s) from redis with error: %v", loginSessionKey, err)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": fmt.Sprintf("Failed to get session from redis: %s", err.Error()),
		})
		return
	}

	// Use token and secret to check user
	userkey, err := misskey.GetUserkey(appSecret, token)
	if err != nil {
		global.Logger.Errorf("Failed to get user key from misskey with error: %v", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": fmt.Sprintf("Failed to get user key from misskey, please contact admin for help."),
		})
		return
	}

	// Else: Valid, delete login session
	global.Redis.Del(context.Background(), loginSessionKey)

	// Check if user is newly registered
	if time.Now().Before(userkey.User.CreatedAt.Add(consts.TIME_NEW_USER_SEND_INVITATION_AFTER)) {
		// Not yet
		ctx.JSON(http.StatusTooEarly, gin.H{
			"error": "Please wait for 24 hours before invite others.",
		})
		return
	}

	adminSession := utils.RandString(32)
	adminSessionKey := fmt.Sprintf(consts.REDIS_KEY_VALID_SESSION, adminSession)

	// Check user
	var user models.User
	err = global.DB.First(&user, "username = ?", userkey.User.Username).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// Create record
			user.Username = userkey.User.Username

			// Insert into database
			err = global.DB.Create(&user).Error
			if err != nil {
				global.Logger.Errorf("Failed to insert user (%s) into database with error: %v", userkey.User.Username, err)
				ctx.JSON(http.StatusInternalServerError, gin.H{
					"error": "Failed to insert user into database.",
				})
				return
			}
		} else {
			global.Logger.Errorf("Failed to get user (%s) data from database with error: %v", userkey.User.Username, err)
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
		"username": userkey.User.Username,
		"name":     userkey.User.Name,
		"avatar":   userkey.User.AvatarUrl,
	})

}
