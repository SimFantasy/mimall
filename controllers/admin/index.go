package admin

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type AdminController struct{}

func (con AdminController) Index(ctx *gin.Context) {

	ctx.HTML(http.StatusOK, "admin/main/index.html", gin.H{})
}

func (con AdminController) Welcome(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "admin/main/welcome.html", gin.H{})
}
