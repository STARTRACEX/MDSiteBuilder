package main

import (
	"md/config"
	"md/model"
	"path"
	"strings"
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
	model.AutoPosts(r.Group("/"))
	model.LangEN(r)
	model.LangZH(r)
	model.LangRU(r)
	r.NoRoute(img)
	r.Run(":" + config.Port)
}
func img(c *gin.Context) {
	pathname, _ := c.GetQuery("pathname")
	pathname = strings.Replace(pathname, config.PostsURL, "", 1)
	filepath, _ := c.GetQuery("filepath")
	e := path.Join(pathname, filepath)
	c.File("." + config.PostsPath + e)
}
