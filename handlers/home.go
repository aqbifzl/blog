package handlers

import (
	"fmt"
	"blog/components"
	"blog/core"

	"github.com/gin-gonic/gin"
)

func GetHome(c *gin.Context) {
	if err := components.Home().Render(c, c.Writer); err != nil {
		fmt.Printf("home rendering: %q", err)
		core.InternalServerErrorPage(c)
		return
	}
}
