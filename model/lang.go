package model

import (
	"md/config"
	"github.com/gin-gonic/gin"
)

type Zh struct{}
type En struct{}
type Ru struct{}

// EN
func LangEN(r *gin.Engine) {
	en := r.Group("/en", func(c *gin.Context) {
		ResetCookie(c, "lang", "en")
		c.Set("lang", "en")
	})
	{
		en.GET("/", func(c *gin.Context) {
			post := ReadMarkdown("/README.md")
			c.HTML(200, "posts.html", gin.H{"Markdown": post})
		})
		AutoPosts(en)
	}
}

// ZH
func LangZH(r *gin.Engine) {
	zh := r.Group("/zh", func(c *gin.Context) {
		ResetCookie(c, "lang", "zh")
		c.Set("lang", "zh")
	})
	{
		zh.GET("/", func(c *gin.Context) {
			post := ReadMarkdown("/README.zh.md")
			c.HTML(200, "posts.html", gin.H{"Markdown": post})
		})
		AutoPosts(zh)
	}
}

// RU
func LangRU(r *gin.Engine) {
	ru := r.Group("/ru", func(c *gin.Context) {
		ResetCookie(c, "lang", "ru")
		c.Set("lang", "ru")
	})
	{
		ru.GET("/", func(c *gin.Context) {
			post := ReadMarkdown("/README.ru.md")
			c.HTML(200, "posts.html", gin.H{"Markdown": post})
		})
		AutoPosts(ru)
	}
}
func AutoPosts(r *gin.RouterGroup) {
	docs := r.Group(config.PostsURL)
	docs.GET("/*url", func(c *gin.Context) { RenderPost(c) })
}
