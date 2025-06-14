package middleware

import (
	"blog/config"
	"blog/core"
	"net/http"
	"path"

	"github.com/gin-gonic/gin"
)

type Intercept404 struct {
	http.ResponseWriter
	statusCode int
}

func (w *Intercept404) Write(b []byte) (int, error) {
	return len(b), nil
}

func (w *Intercept404) WriteHeader(statusCode int) {
	w.statusCode = statusCode
}

func SPAMiddleware(conf *config.AppConfig) func(c *gin.Context) {
	fs := http.StripPrefix("/static", http.FileServer(http.Dir("./static")))

	return func(c *gin.Context) {
		wt := &Intercept404{ResponseWriter: c.Writer}

		original := c.Request.URL.Path
		c.Request.URL.Path = path.Join("/static", original)

		fs.ServeHTTP(wt, c.Request)

		if wt.statusCode == http.StatusNotFound {
			c.Header("Content-Type", "text/html")
			core.NotFoundPage(c, conf)
			return
		}

		fs.ServeHTTP(c.Writer, c.Request)
	}
}
