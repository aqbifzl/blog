package handlers

import (
	"blog/config"
	"blog/middleware"

	"github.com/gin-gonic/gin"
)

func SetupServer(conf *config.AppConfig) *gin.Engine {
	r := gin.Default()

	SetupHomeHandlers(r, conf)
	SetupBlogHandlers(r, conf)
	SetupContactHandler(r, conf)

	r.NoRoute(middleware.SPAMiddleware(conf))

	return r
}
