package main

import (
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"md/config"
	"md/model"
	"path"
	"strings"
)

func main() {
	if !config.Dev {
		gin.SetMode(gin.ReleaseMode)
		gin.DefaultWriter = ioutil.Discard
	}
	r := gin.Default()

	r.LoadHTMLGlob(config.HtmlGlobPath)
	r.Static(config.StaticPath, "."+config.StaticPath)
	r.Static(config.PostsPath, "."+config.PostsPath)
	r.StaticFile("/favicon.ico", "./favicon.ico")
	r.StaticFile("/robots.txt", "./robots.txt")
	r.StaticFile("/README.md", "README.md")
	r.GET("/", func(c *gin.Context) {
		post := model.ReadMarkdown("/README.md")
		c.HTML(200, "posts.html", gin.H{"Markdown": post})
	})
	model.AutoPosts(r.Group("/"))
	model.LangEN(r)
	model.LangZH(r)
	model.LangRU(r)
	r.GET("/img", img)
	r.Run(":" + config.Port)
}
func img(c *gin.Context) {
	pathname, _ := c.GetQuery("pathname")
	filepath, _ := c.GetQuery("filepath")
	if strings.Contains(pathname, config.PostsURL) {
		pathname = strings.Replace(pathname, config.PostsURL, "", 1)
		e := path.Join(path.Dir(pathname), filepath)
		c.File("." + config.PostsPath + e)
	} else {
		e := path.Join(path.Dir(pathname), filepath)
		c.File("." + e)
	}
}
