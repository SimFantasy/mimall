package middlewares

import (
	"encoding/json"
	"mimall/models"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func AdminAuthMiddleware(ctx *gin.Context) {
	s := sessions.Default(ctx)
	u := s.Get("userinfo")
	uStr, ok := u.(string)
	if ok {
		var manager []models.Manager
		json.Unmarshal([]byte(uStr), &manager)
		ctx.JSON(http.StatusOK, gin.H{
			"user": manager[0].Username,
		})
	}
}
