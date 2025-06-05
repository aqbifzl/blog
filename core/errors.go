package core

import (
	"net/http"
	"blog/components"

	"github.com/gin-gonic/gin"
)

var (
	InternalServerErrorTitle = "Coś poszło nie tak"
	InternalServerErrorDesc  = "Bardzo złe rzeczy stały się na serwerze"
	NotFoundTitle            = "Nic tu nie ma"
	NotFoundDesc             = "I jeśli nic nie popsułem to pewnie nie będzie"
	BadRequest               = "Nieprawidłowe zapytanie"
)

func BadRequestPage(c *gin.Context, desc string) {
	c.Status(http.StatusBadRequest)
	if err := components.ErrorPage(NotFoundTitle, desc).Render(c, c.Writer); err != nil {
		InternalServerErrorPage(c)
	}
}

func NotFoundPage(c *gin.Context) {
	c.Status(http.StatusNotFound)
	if err := components.ErrorPage(NotFoundTitle, NotFoundDesc).Render(c, c.Writer); err != nil {
		InternalServerErrorPage(c)
	}
}

func InternalServerErrorPage(c *gin.Context) {
	c.Status(http.StatusInternalServerError)
	_ = components.ErrorPage(InternalServerErrorTitle, InternalServerErrorDesc).Render(c, c.Writer)
}
