index
index summary
2020
2022
x
## index page of zh
```go


func RenderDir(router *gin.Engine, dir string) {
	_g := router.Group(dir+"/:g", func(ctx *gin.Context) { ctx.Set("g", ctx.Param("g")); ctx.Set("lang", dir) })
	{
		/* lv1
		default
		*/
		_g.GET("/", func(ctx *gin.Context) {
			lang := Sget(ctx, "lang")
			g, _ := ctx.Get("g")
			// 如果/posts/:lang/def/:g不存在，目录为posts/:lang/:g
			path := config.PostsPath + lang + config.PostDefaultdPath
			if !IsExist("." + path) {
				path = config.PostsPath + lang
			}
			path = path + "/" + g.(string)
			fmt.Println(">", path, "<")
			if !IsExist("." + path + ".md") {
				path += "/index"
			}
			post := MarkdownPost(path)
			ctx.HTML(200, "posts.html", gin.H{"Markdown": post})
		})
		_c := _g.Group("/:c", func(ctx *gin.Context) { ctx.Set("c", ctx.Param("c")) })
		{
			/* lv2 */
			_c.GET("/", func(ctx *gin.Context) {
				lang := Sget(ctx, "lang")
				fmt.Println("lang::", lang)
				g, _ := ctx.Get("g")
				c, _ := ctx.Get("c")
				sg := g.(string)
				sc := c.(string)
				path := config.PostsPath + lang + "/" + sg + "/" + sc
				if !IsExist("." + path + ".md") {
					path = config.PostsPath + lang + "/" + sg + "/" + sc + "/index"
				}
				fmt.Println(path)
				post := MarkdownPost(path)
				ctx.HTML(200, "posts.html", gin.H{"Markdown": post})
			})
			{
				/* lv3 */
				_c.GET("/:p", func(ctx *gin.Context) {
					lang := Sget(ctx, "lang")
					c, _ := ctx.Get("c")
					g, _ := ctx.Get("g")
					p := ctx.Param("p")
					sg := g.(string)
					sc := c.(string)
					path := config.PostsPath + lang + "/" + sg + "/" + sc + "/" + p
					post := MarkdownPost(path)
					ctx.HTML(200, "posts.html", gin.H{"Markdown": post})
				})
				/* lv3 */
			}
			/* lv2 */
		}
		/* lv1 */
	}
}

```