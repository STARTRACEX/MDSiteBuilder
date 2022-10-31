package model

import (
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
		AutoPosts(ru)
	}
}
func AutoPosts(r *gin.RouterGroup) {
	r.GET("/docs", func(c *gin.Context) {
		RenderPost(c)
	})
	r.GET("/docs/*url", func(c *gin.Context) {
		RenderPost(c)
	})
}
