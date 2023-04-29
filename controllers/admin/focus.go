package admin

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type FocusController struct{}

func (con FocusController) Index(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "admin/focus/index.html", gin.H{})
}

func (con FocusController) Add(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "admin/focus/add.html", gin.H{})
}

func (con FocusController) DoAdd(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "admin/focus/do_add.html", gin.H{})
}

func (con FocusController) Edit(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "admin/focus/edit.html", gin.H{})
}

func (con FocusController) DoEdit(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "admin/focus/do_edit.html", gin.H{})
}

func (con FocusController) Delete(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "delete success page",
	})
}
