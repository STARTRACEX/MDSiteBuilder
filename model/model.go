package model

import (
	"github.com/gin-gonic/gin"
	"github.com/russross/blackfriday"
	"html/template"
	"io/ioutil"
	"md/config"
	"os"
	"strings"
)

type Post struct {
	Title      string
	Date       string
	Update     string
	Summary    string
	Author     template.HTML
	Body       template.HTML
	OriginFile string // Origin File Path
}
type Cata struct {
	Name  string
	Super string
	Level string
}

func RenderMarkdown(mdFilePath string) Post {
	source := "" + config.PostsPath + "/" + mdFilePath + ".md"
	if !IsExist(source) {
		source = "" + config.PostsPath + "/" + mdFilePath + "index.md"
	}
	if !IsExist(source) {
		source = "" + config.PostsPath + "/" + mdFilePath + "/index.md"
	}
	return ReadMarkdown(source)
}

func RenderPost(c *gin.Context) {
	var dir string = "zh"
	l, e := c.Get("lang")
	if e {
		dir = l.(string)
	}
	cata := RenderCatalog("." + config.PostsPath + "/" + dir + "/cata.txt")
	url := dir + "/" + c.Param("url")
	post := RenderMarkdown(url)
	c.HTML(200, "posts.html", gin.H{"Markdown": post, "Catalog": cata})
}

func RenderCatalog(DirectoryFilePath string) []Cata {
	data, _ := ioutil.ReadFile(DirectoryFilePath)
	lines := strings.Split(string(data), "\n")
	var C []Cata
	for _, line := range lines {
		line = strings.ReplaceAll(line, "\r", "")
		s := strings.Split(line, "\\")
		if len(s) < 2 {
			break
		}
		S := Cata{
			Super: s[2],
			Name:  s[1],
			Level: s[0],
		}
		C = append(C, S)
	}
	return C
}
func ReadMarkdown(source string) Post {
	var body, summary, author string
	var title, date, update string = "Undefined", "Unknow", "Unknow"
	fileread, err := ioutil.ReadFile("." + source)
	if err != nil {
		return Post{
			Title:      title,
			Summary:    summary,
			Date:       date,
			Update:     update,
			Author:     template.HTML(author),
			Body:       template.HTML(body),
			OriginFile: source,
		}
	}
	lines := strings.Split(string(fileread), "\n")
	if len(lines) > 0 {
		title = strings.ReplaceAll(strings.ReplaceAll(string(lines[0]), "\r", ""), "# ", "")
	}
	if len(lines) > 1 {
		summary = string(lines[1])
	}
	if len(lines) > 2 {
		date = string(lines[2])
	}
	if len(lines) > 3 {
		update = string(lines[3])
	}
	if len(lines) > 4 {
		author = string(blackfriday.MarkdownCommon([]byte(lines[4])))
	}
	if len(lines) >= 5 {
		body = string(blackfriday.MarkdownCommon([]byte(lines[0]))) + string(blackfriday.MarkdownCommon([]byte(strings.Join(lines[5:], ""))))
	}
	return Post{
		Title:      title,
		Summary:    summary,
		Date:       date,
		Update:     update,
		Author:     template.HTML(author),
		Body:       template.HTML(body),
		OriginFile: source}
}
func IsExist(path string) bool {
	_, err := os.Stat(path)
	return err == nil || os.IsExist(err)
}
