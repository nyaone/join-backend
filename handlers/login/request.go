package login

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"join-nyaone/config"
	"join-nyaone/consts"
	"join-nyaone/global"
	"join-nyaone/misskey"
	"join-nyaone/utils"
	"net/http"
	"strings"
	"time"
)

func Request(ctx *gin.Context) {
	username := strings.ToLower(ctx.Param("username"))

	// Check if in redis
	sessionKey := fmt.Sprintf(consts.REDIS_KEY_LOGIN_REQUEST, username)
	if exist, err := global.Redis.Exists(context.Background(), sessionKey).Result(); err != nil {
		global.Logger.Errorf("Failed to check session key (%s) from redis with error: %v", sessionKey, err)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": fmt.Sprintf("Failed to check session from redis: %s", err.Error()),
		})
		return
	} else if exist > 0 {
		ctx.JSON(http.StatusTooManyRequests, gin.H{
			"error": "Session already exist.",
		})
		return
	}

	// Get user info
	userInfo, err := misskey.GetUser(username)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	// Check if is in cool-down period
	if time.Now().Before(userInfo.CreatedAt.Add(consts.TIME_NEW_USER_SEND_INVITATION_AFTER)) {
		// Not yet
		ctx.JSON(http.StatusTooEarly, gin.H{
			"error": "Please wait for 24 hours before invite others.",
		})
		return
	}

	// Prepare login secret
	secret := utils.RandString(8)

	// Save secret to redis
	err = global.Redis.Set(context.Background(), sessionKey, secret, consts.TIME_LOGIN_REQUEST_VALID).Err()
	if err != nil {
		global.Logger.Errorf("Failed to save session secret (%s) into redis with error: %v", sessionKey, err)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	// Send message
	token := fmt.Sprintf("%s-%s", username, secret)
	link := fmt.Sprintf("%s/login/%s", config.Config.FrontendUri, token)
	text := fmt.Sprintf(consts.MESSAGE_TEMPLATE, link)
	sendMsgRes, err := misskey.SendMessage(userInfo.ID, text)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	// Return true
	ctx.JSON(http.StatusOK, gin.H{
		"ok":             true,
		"messaging_link": fmt.Sprintf("%s/my/messaging/%s", config.Config.Misskey.Instance, sendMsgRes.User.Username),
	})

}
