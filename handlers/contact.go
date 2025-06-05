package handlers

import (
	"blog/components"
	"blog/config"
	"blog/core"
	"fmt"

	"github.com/gin-gonic/gin"
)

type ContactHandler struct {
	conf *config.AppConfig
}

func SetupContactHandler(r *gin.Engine, conf *config.AppConfig) *ContactHandler {
	contact_handler := ContactHandler{conf: conf}

	r.GET("/contact", contact_handler.GetContact)

	return &contact_handler
}

func (h ContactHandler) GetContact(c *gin.Context) {
	if err := components.Contact(h.conf).Render(c, c.Writer); err != nil {
		fmt.Printf("contact rendering: %q", err)
		core.InternalServerErrorPage(c, h.conf)
		return
	}
}
