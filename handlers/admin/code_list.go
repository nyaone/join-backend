package admin

import (
	"github.com/gin-gonic/gin"
	"join-nyaone/global"
	"join-nyaone/handlers/invitee"
	"join-nyaone/models"
	"net/http"
)

func CodeList(ctx *gin.Context) {
	userId, exist := ctx.Get("UID")
	if !exist {
		global.Logger.Errorf("No UID provided by middleware")
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "Middleware not working properly",
		})
		return
	}

	var inviteCodes []models.Code
	responseCodes := []CodeResponse{} // So it won't be null
	global.DB.Find(&inviteCodes, "created_by_user_id = ?", userId.(uint))

	for _, inviteCode := range inviteCodes {

		isValid, inviteCount := invitee.CheckInviteCodeValid(&inviteCode)
		responseCodes = append(responseCodes, CodeResponse{
			Code:        inviteCode.Code.Code.String(),
			CodeProps:   inviteCode.CodeProps,
			InviteCount: inviteCount,
			IsValid:     isValid,
		})
	}

	ctx.JSON(http.StatusOK, responseCodes)

}
