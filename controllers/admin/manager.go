package admin

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ManagerController struct{}

func (con ManagerController) Index(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "admin/manager/index.html", gin.H{})
}

func (con ManagerController) Add(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "admin/manager/add.html", gin.H{})
}

func (con ManagerController) DoAdd(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "admin/manager/do_add.html", gin.H{})
}

func (con ManagerController) Edit(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "admin/manager/edit.html", gin.H{})
}

func (con ManagerController) DoEdit(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "admin/manager/do_edit.html", gin.H{})
}

func (con ManagerController) Delete(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "delete success page",
	})
}
