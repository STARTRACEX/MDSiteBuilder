package main

import (
	"md/config"
	"md/model"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.LoadHTMLGlob(config.HtmlGlobPath)
	r.Static(config.StaticPath, "."+config.StaticPath)
	r.Static(config.PostsPath, "."+config.PostsPath)
	r.StaticFile("/favicon.ico", "./favicon.ico")
	r.StaticFile("/README.md", "README.md")
	r.GET("/", func(c *gin.Context) {
		post := model.ReadMarkdown("/README.md")
		c.HTML(200, "posts.html", gin.H{"Markdown": post})
	})
	g := r.Group("/")
	model.AutoPosts(g)
	model.LangEN(r)
	model.LangZH(r)
	model.LangRU(r)
	r.Run(":" + config.Port)
}
