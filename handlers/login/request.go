package login

import (
	"context"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"join-nyaone/config"
	"join-nyaone/consts"
	"join-nyaone/global"
	"join-nyaone/misskey"
	"join-nyaone/models"
	"join-nyaone/types"
	"net/http"
)

func Request(ctx *gin.Context) {

	// Check current application. If not exist, create one.
	var app models.Application
	if err := global.DB.First(&app, "instance_uri = ? AND frontend_uri = ?", config.Config.Misskey.Instance, config.Config.FrontendUri).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// No matching application, create one

			// Init meta information
			app.InstanceUri = config.Config.Misskey.Instance
			app.FrontendUri = config.Config.FrontendUri
			app.ApplicationPublic = types.ApplicationPublic{
				Name:        "Join NyaOne",
				Description: "喵窝邀请管理系统",
				Permission:  []string{},
				CallbackURL: fmt.Sprintf("%s/login", config.Config.FrontendUri),
			}

			// Get specified
			newApp, err := misskey.CreateApplication(&app.ApplicationPublic)
			if err != nil {
				global.Logger.Errorf("Failed to create auth application with error: %v", err)
				ctx.JSON(http.StatusInternalServerError, gin.H{
					"error": "Failed to create auth application, please contact admin for help.",
				})
				return
			}
			app.ApplicationPrivate = newApp.ApplicationPrivate

			// Save into database
			if err = global.DB.Save(&app).Error; err != nil {
				global.Logger.Errorf("Failed to save auth application %v into database with error: %v", app, err)
				ctx.JSON(http.StatusInternalServerError, gin.H{
					"error": "Failed to create auth application, please contact admin for help.",
				})
				return
			}
		} else {
			// Other errors
			global.Logger.Errorf("Failed to check auth application from database with error: %v", err)
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error": "Failed to check auth application, please contact admin for help.",
			})
			return
		}
	}

	// Create auth request
	sess, err := misskey.GenerateAuthSession(app.Secret)
	if err != nil {
		global.Logger.Errorf("Failed to generate auth session with error: %v", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to generate auth session, please contact admin for help.",
		})
		return
	}

	// Save into redis
	sessionKey := fmt.Sprintf(consts.REDIS_KEY_LOGIN_REQUEST, sess.Token)
	err = global.Redis.Set(context.Background(), sessionKey, app.Secret, consts.TIME_LOGIN_REQUEST_VALID).Err()

	// Return true
	ctx.JSON(http.StatusOK, gin.H{
		"ok":  true,
		"url": sess.Url,
	})

}
