package handlers

import (
	"blog/components"
	"blog/config"
	"blog/core"
	"blog/services"
	"blog/utils"
	"errors"
	"fmt"
	"math"
	"net/http"
	"os"
	"path/filepath"
	"sort"
	"strconv"

	"github.com/gin-gonic/gin"
)

type BlogHandler struct {
	BlogService *services.BlogService
	conf        *config.AppConfig
}

func SetupBlogHandlers(r *gin.Engine, conf *config.AppConfig) {
	blog_handler := BlogHandler{
		BlogService: services.NewBlogService(),
		conf:        conf,
	}

	r.GET("/blog", blog_handler.GetPosts)
	r.GET("/blog/:page", blog_handler.GetPostsPage)
	r.GET("/post/:name", blog_handler.GetPost)
}

func (h BlogHandler) handle_posts(c *gin.Context, page int) {
	if page < 1 {
		c.Redirect(http.StatusMovedPermanently, "/blog")
		return
	}

	posts, err := h.BlogService.GetPostsMetadata(page)
	if err != nil {
		fmt.Printf("getting posts: %q", err)

		core.InternalServerErrorPage(c, h.conf)
		return
	}

	if len(posts) < 1 {
		if err := components.ErrorPage("Nie ma postów", "Proszę o cierpliwość", h.conf).Render(c, c.Writer); err != nil {
			core.InternalServerErrorPage(c, h.conf)
		}
		return
	}

	pages := int(math.Ceil(float64(len(posts)) / config.POSTS_PER_PAGE))

	if page > pages {
		core.NotFoundPage(c, h.conf)
		return
	}

	sort.Slice(posts, func(i, j int) bool {
		return posts[i].Date.After(posts[j].Date)
	})

	first_post := config.POSTS_PER_PAGE * (page - 1)
	last := first_post + config.POSTS_PER_PAGE
	posts = posts[first_post:min(last, len(posts))]

	if err := components.Blog(posts, page, pages, utils.GetPageList(page, pages), h.conf).Render(c, c.Writer); err != nil {
		fmt.Printf("rendering blog: %q", err)
		core.InternalServerErrorPage(c, h.conf)
		return
	}
}

func (h BlogHandler) GetPostsPage(c *gin.Context) {
	page_int := 1
	page := c.Param("page")
	if len(page) != 0 {
		if conv, err := strconv.ParseInt(page, 10, 32); err == nil {
			page_int = int(conv)
		} else {
			c.Redirect(http.StatusPermanentRedirect, "/blog")
			return
		}
	}

	h.handle_posts(c, page_int)
}

func (h BlogHandler) GetPosts(c *gin.Context) {
	h.handle_posts(c, 1)
}

func (h BlogHandler) GetPost(c *gin.Context) {
	name := c.Param("name")
	if len(name) < 1 || !utils.IsFilenameSafe(name) {
		core.BadRequestPage(c, "Podejrzany url", h.conf)
		return
	}

	path := filepath.Join("./posts", filepath.Clean(name))
	if _, err := os.Stat(path); err != nil {
		if errors.Is(err, os.ErrNotExist) {
			core.NotFoundPage(c, h.conf)
		} else {
			fmt.Printf("checking if path %q exists: %q", path, err)
			core.InternalServerErrorPage(c, h.conf)
		}
		return
	}

	post, err := h.BlogService.GetBlogPost(path)
	if err != nil {
		fmt.Printf("getting blog post: %q", err)
		core.InternalServerErrorPage(c, h.conf)
		return
	}

	if err := components.Post(post, h.conf).Render(c, c.Writer); err != nil {
		fmt.Printf("rendering post: %q", err)
		core.InternalServerErrorPage(c, h.conf)
		return
	}
}
