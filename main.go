package main

import (
	"github.com/gin-gonic/gin"
	"md/config"
	"md/model"
	"strings"
)

func main() {
	if !config.Dev {
		gin.SetMode(gin.ReleaseMode)
	}
	r := gin.New()
	r.Use(gin.Recovery())
	r.LoadHTMLGlob(config.HtmlGlobPath)
	/* Static file */
	r.Static(config.StaticPath, "."+config.StaticPath)
	if !strings.HasPrefix(config.PostsPath, config.StaticPath) {
		r.Static(config.PostsPath, "."+config.PostsPath)
	}
	r.StaticFile("/favicon.ico", "."+config.StaticPath+"/favicon.ico")
	r.StaticFile("/robots.txt", "."+config.StaticPath+"/robots.txt")
	r.StaticFile("/README.md", "./README.md")
	r.StaticFile("/README.zh.md", "./README.zh.md")
	r.StaticFile("/README.ru.md", "./README.ru.md")

	/* root */
	r.GET("/", func(c *gin.Context) {
		post := model.ReadMarkdown("/README.md")
		c.HTML(200, "posts.html", gin.H{"Markdown": post})
	})
	model.AutoPosts(r.Group("/"))
	/* main */
	model.I18ninit(r)
	/* img */
	r.GET("/img", model.Img)
	r.Run(":" + config.Port)
}
