package handlers

import (
	"blog/config"
	"blog/core"

	"github.com/gin-gonic/gin"
)

func SetupServer(conf *config.AppConfig) *gin.Engine {
	r := gin.Default()

	r.Static("/static", "./static")

	r.GET("/", GetHome)
	SetupBlogHandlers(r)
	SetupContactHandler(r, conf)
	r.NoRoute(func(c *gin.Context) {
		core.NotFoundPage(c)
	})

	return r
}
