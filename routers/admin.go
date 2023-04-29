package routers

import (
	"mimall/controllers/admin"
	"mimall/middlewares"

	"github.com/gin-gonic/gin"
)

func AdminRouters(r *gin.Engine) {
	adminRouter := r.Group("/admin", middlewares.AdminAuthMiddleware)
	{
		adminRouter.GET("/", admin.AdminController{}.Index)
		adminRouter.GET("/welcome", admin.AdminController{}.Welcome)
		adminRouter.GET("/login", admin.LoginContrller{}.Index)
		adminRouter.GET("/captcha", admin.LoginContrller{}.Captcha)
		adminRouter.POST("/do_login", admin.LoginContrller{}.DoLogin)
		adminRouter.GET("/logout", admin.LoginContrller{}.Logout)

		adminRouter.GET("/manager", admin.ManagerController{}.Index)
		adminRouter.GET("/manager/add", admin.ManagerController{}.Add)
		adminRouter.POST("/manager/do_add", admin.ManagerController{}.DoAdd)
		adminRouter.GET("/manager/edit", admin.ManagerController{}.Edit)
		adminRouter.PUT("/manager/do_edit", admin.ManagerController{}.DoEdit)
		adminRouter.DELETE("/manager/delete", admin.ManagerController{}.Delete)

		adminRouter.GET("/focus", admin.FocusController{}.Index)
		adminRouter.GET("/focus/add", admin.FocusController{}.Add)
		adminRouter.POST("/focus/do_add", admin.FocusController{}.DoAdd)
		adminRouter.GET("/focus/edit", admin.FocusController{}.Edit)
		adminRouter.PUT("/focus/do_edit", admin.FocusController{}.DoEdit)
		adminRouter.DELETE("/focus/delete", admin.FocusController{}.Delete)
	}
}
