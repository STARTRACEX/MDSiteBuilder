package main

import (
	"html/template"
	"io/ioutil"
	"md/config"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/russross/blackfriday"
)

var G, C, P string

func main() {
	router := gin.Default()
	router.LoadHTMLGlob(config.HtmlGlobPath)
	router.Static(config.StaticPath, "."+config.StaticPath)
	router.Static(config.PostsPath, "."+config.PostsPath)
	router.StaticFile("/favicon.ico", "./favicon.ico")
	router.StaticFile("/README.md", "README.md")
	// router.GET("/catalog",func(ctx *gin.Context) {
	// 	ctx.HTML(200, "test.html", gin.H{"catalog": WalkFiles})
	// })
	router.GET("/", func(ctx *gin.Context) {
		post := MarkdownPost("/README")
		ctx.HTML(200, "posts.html", gin.H{"Markdown": post})
	})
	
	for _, v := range config.MDGL {
		RenderDir(router, v)
	}

	router.Run(":" + config.Port)
}

type Post struct {
	Title      string
	Date       string
	Update     string
	Summary    string
	Auther     string
	Body       template.HTML
	OriginFile string
}

func MarkdownPost(filepath string) Post {
	fileread, _ := ioutil.ReadFile("." + filepath + ".md")
	lines := strings.Split(string(fileread), "\n")
	var (
		summary, body, auther string
		title, date, update   string = "Undefined", "Unknow", "Unknow"
	)
	if len(lines) >= 5 {
		title = strings.Replace(string(lines[0]), "\r", "", 1)
		summary = string(lines[1])
		date = string(lines[2])
		update = string(lines[3])
		auther = string(lines[4])
		body = strings.Join(lines[5:], "")
	}
	body = string(blackfriday.MarkdownCommon([]byte(body)))
	return Post{
		Title:      title,
		Date:       date,
		Update:     update,
		Summary:    summary,
		Auther:     auther,
		Body:       template.HTML(body),
		OriginFile: filepath + ".md",
	}
}

func RenderDir(router *gin.Engine, dir string) {
	_g := router.Group(dir+"/:g", func(ctx *gin.Context) {
		ctx.Set("g", ctx.Param("g"))
		ctx.Set("lang", dir)

	})
	{
		/* lv1 */
		_g.GET("/", func(ctx *gin.Context) {
			lang := Sget(ctx, "lang")
			g := Sget(ctx, "g")
			// Relocate File
			path := config.PostsPath + lang + config.PostDefaultdPath
			if !IsExist("." + path) {
				path = config.PostsPath + lang
			}
			path = path + "/" + g
			if !IsExist("." + path + ".md") {
				path += "/index"
			}
			post := MarkdownPost(path)
			ctx.HTML(200, "posts.html", gin.H{"Markdown": post, "Lang": ReturnLang(lang)})
		})
		_c := _g.Group("/:c", func(ctx *gin.Context) { ctx.Set("c", ctx.Param("c")) })
		{
			/* lv2 */
			_c.GET("/", func(ctx *gin.Context) {
				lang := Sget(ctx, "lang")
				g := Sget(ctx, "g")
				c := Sget(ctx, "c")
				path := config.PostsPath + lang + "/" + g + "/" + c
				if !IsExist("." + path + ".md") {
					path += "/index"
				}
				post := MarkdownPost(path)
				ctx.HTML(200, "posts.html", gin.H{"Markdown": post, "Lang": ReturnLang(lang)})
			})
			{
				/* lv3 */
				_c.GET("/:p", func(ctx *gin.Context) {
					lang := Sget(ctx, "lang")
					g := Sget(ctx, "g")
					c := Sget(ctx, "c")
					p := ctx.Param("p")
					path := config.PostsPath + lang + "/" + g + "/" + c + "/" + p
					post := MarkdownPost(path)
					ctx.HTML(200, "posts.html", gin.H{"Markdown": post, "Lang": ReturnLang(lang)})
				})
				/* lv3 */
			}
			/* lv2 */
		}
		/* lv1 */
	}
}

func Sget(ctx *gin.Context, name string) string {
	v, e := ctx.Get(name)
	if e || v != nil {
		return v.(string)
	} else {
		return ""
	}
}
func ReturnLang(s string) string {
	if s == "" {
		return "global"
	}
	return strings.Replace(s, "/", "", 1)
}
func IsExist(path string) bool {
	_, err := os.Stat(path)
	return err == nil || os.IsExist(err)
}