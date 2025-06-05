package core

import (
	"blog/components"
	"blog/config"
	"net/http"

	"github.com/gin-gonic/gin"
)

var (
	InternalServerErrorTitle = "Coś poszło nie tak"
	InternalServerErrorDesc  = "Bardzo złe rzeczy stały się na serwerze"
	NotFoundTitle            = "Nic tu nie ma"
	NotFoundDesc             = "I jeśli nic nie popsułem to pewnie nie będzie"
	BadRequest               = "Nieprawidłowe zapytanie"
)

func BadRequestPage(c *gin.Context, desc string, conf *config.AppConfig) {
	c.Status(http.StatusBadRequest)
	if err := components.ErrorPage(NotFoundTitle, desc, conf).Render(c, c.Writer); err != nil {
		InternalServerErrorPage(c, conf)
	}
}

func NotFoundPage(c *gin.Context, conf *config.AppConfig) {
	c.Status(http.StatusNotFound)
	if err := components.ErrorPage(NotFoundTitle, NotFoundDesc, conf).Render(c, c.Writer); err != nil {
		InternalServerErrorPage(c, conf)
	}
}

func InternalServerErrorPage(c *gin.Context, conf *config.AppConfig) {
	c.Status(http.StatusInternalServerError)
	_ = components.ErrorPage(InternalServerErrorTitle, InternalServerErrorDesc, conf).Render(c, c.Writer)
}
