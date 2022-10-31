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
	// r.StaticFile("/README.md", "README.md")
	// r.GET("/", func(ctx *gin.Context) {
	// 	post := model.MarkdownPost("/README")
	// 	ctx.HTML(200, "posts.html", gin.H{"Markdown": post})
	// })
	r.GET("/docs", model.RenderPost)
	r.GET("/docs/*url", model.RenderPost)
	r.Run(":" + config.Port)
}

