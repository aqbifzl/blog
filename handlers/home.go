package handlers

import (
	"blog/components"
	"blog/config"
	"blog/core"
	"fmt"

	"github.com/gin-gonic/gin"
)

type HomeHandler struct {
	conf *config.AppConfig
}

func SetupHomeHandlers(r *gin.Engine, conf *config.AppConfig) {
	home_handler := HomeHandler{conf: conf}

	r.GET("/", home_handler.GetHome)
}

func (h HomeHandler) GetHome(c *gin.Context) {
	if err := components.Home(h.conf).Render(c, c.Writer); err != nil {
		fmt.Printf("home rendering: %q", err)
		core.InternalServerErrorPage(c, h.conf)
		return
	}
}
