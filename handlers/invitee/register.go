package invitee

import (
	"github.com/gin-gonic/gin"
	"join-nyaone/config"
	"join-nyaone/global"
	"join-nyaone/misskey"
	"join-nyaone/models"
	"join-nyaone/types"
	"net/http"
)

type RegisterRequest struct {
	InviteCode string `json:"invite_code"` // UUID in string
	Username   string `json:"username"`
	Password   string `json:"password"`
}

func Register(ctx *gin.Context) {
	var req RegisterRequest
	err := ctx.BindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Cannot parse request data",
		})
		return
	}

	// Check invite code
	isInviteCodeValid, errCode, targetInviteCode, err := CheckInviteCodeStatus(req.InviteCode)
	if err != nil {
		ctx.JSON(errCode, gin.H{
			"error": err.Error(),
		})
		return
	} else if !isInviteCodeValid {
		ctx.JSON(http.StatusOK, gin.H{
			"error": "Sorry, this invite code is invalid.",
		})
		return
	}

	// Check username
	isUsernameValid, err := misskey.CheckUsername(req.Username)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	} else if !isUsernameValid {
		ctx.JSON(http.StatusOK, gin.H{
			"error": "Sorry, this username is invalid.",
		})
		return
	}

	// Do register!
	newAccount, err := misskey.CreateAccount(req.Username, req.Password)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "Sorry, we cannot create your account at this moment, please try again later.",
		})
		return
	}

	// Save into database
	invitee := models.User{
		User: types.User{
			Username:        newAccount.Username,
			InvitedByCodeID: targetInviteCode.ID,
			InvitedByUserID: targetInviteCode.CreatedByUserID,
		},
	}

	global.DB.Create(&invitee)

	ctx.JSON(http.StatusOK, gin.H{
		"ok":       true,
		"username": newAccount.Username,
		"instance": config.Config.Misskey.Instance,
	})

}
