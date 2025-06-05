package handlers

import (
	"blog/config"
	"blog/core"

	"github.com/gin-gonic/gin"
)

func SetupServer(conf *config.AppConfig) *gin.Engine {
	r := gin.Default()

	r.Static("/static", "./static")

	SetupHomeHandlers(r, conf)
	SetupBlogHandlers(r, conf)
	SetupContactHandler(r, conf)
	r.NoRoute(func(c *gin.Context) {
		core.NotFoundPage(c, conf)
	})

	return r
}
