package admin

import (
	"encoding/json"
	"fmt"
	"mimall/models"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

type LoginContrller struct {
	BaseController
}

func (con LoginContrller) Index(ctx *gin.Context) {
	// fmt.Println(models.Md5("123456"))
	ctx.HTML(http.StatusOK, "admin/login/index.html", gin.H{})
}

func (con LoginContrller) DoLogin(ctx *gin.Context) {

	username := ctx.PostForm("username")
	password := ctx.PostForm("password")
	fmt.Printf("username:%v password:%v pswType:%T \n", username, password, password)

	value := ctx.PostForm("verifyValue")
	id := ctx.PostForm("captcha_id")

	if flag := models.VerifyCaptcha(id, value); flag {
		// checkout username and password
		managerInfo := []models.Manager{}
		// md5 password
		password = models.Md5(password)

		models.DB.Where("username=? AND password=?", username, password).Find(&managerInfo)

		if len(managerInfo) > 0 {
			session := sessions.Default(ctx)
			jsonSlice, err := json.Marshal(managerInfo)
			if err != nil {
				fmt.Printf("Marshal err: %v\n", err)
			}
			session.Set("userinfo", string(jsonSlice))
			session.Save()
			con.Success(ctx, "登录成功 ~", "/admin")
		} else {
			con.Error(ctx, "用户名或密码错误 ~", "/admin/login")
		}

	} else {
		con.Error(ctx, "验证码错误 ~", "/admin/login")
	}

}

func (con LoginContrller) Captcha(ctx *gin.Context) {
	id, b64s, err := models.MakeCaptcha()
	if err != nil {
		fmt.Printf("Captcha err: %v\n", err)
	}
	ctx.JSON(http.StatusOK, gin.H{
		"captcha_id":    id,
		"captcha_image": b64s,
	})
}

func (con LoginContrller) Logout(ctx *gin.Context) {
	// delete session
	session := sessions.Default(ctx)
	session.Delete("userinfo")
	session.Save()
	con.Success(ctx, "退出成功 ~", "/admin/login")
}
