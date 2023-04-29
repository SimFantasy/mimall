package main

import (
	"mimall/routers"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// load templates and static files
	r.LoadHTMLGlob("./templates/**/**/*")
	r.Static("/static", "./static")

	// session middleware
	store := cookie.NewStore([]byte("simmimall"))
	r.Use(sessions.Sessions("mimall", store))

	// routers
	routers.AdminRouters(r)

	r.Run(":9900")
}
