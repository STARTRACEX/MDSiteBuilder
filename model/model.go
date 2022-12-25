package model

import (
	"github.com/gin-gonic/gin"
	"github.com/goccy/go-yaml"
	"github.com/russross/blackfriday"
	"html/template"
	"io/ioutil"
	"md/config"
	"os"
	"path"
	"regexp"
	"strings"
)

// md文件的元数据结构
type baseconfig struct {
	Title   string   `yaml:"title,omitempty"`
	Author  string   `yaml:"author,omitempty"`
	Authors []string `yaml:"authors,omitempty"`
	Summary string   `yaml:"summary,omitempty"`
	Date    string   `yaml:"date,omitempty"`
	Update  string   `yaml:"update,omitempty"`
}

// 渲染文档的元数据结构
type Post struct {
	Title      string
	Date       string
	Update     string
	Summary    string
	Author     template.HTML
	Body       template.HTML
	OriginFile string // Origin File Path
}

// 目录文件数据结构
type Cata struct {
	Name  string
	Super string
	Level string
}

func RenderMarkdown(c *gin.Context, mdFilePath string) Post {
	if path.Ext(mdFilePath) != ".md" || path.Ext(mdFilePath) != "" {
		c.Abort()
	}
	source := config.PostsPath + "/" + mdFilePath + ".md"
	if !IsExist("." + source) {
		source = config.PostsPath + "/" + mdFilePath + "index.md"
		if !IsExist("." + source) {
			source = config.PostsPath + "/" + mdFilePath + "/index.md"
			if !IsExist("." + source) {
				source = config.PostsPath + "/" + mdFilePath
			}
		}
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
	url := dir + c.Param("url")
	post := RenderMarkdown(c, url)
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
	meta, metalen := getMeta(string(fileread))
	var config baseconfig
	yaml.Unmarshal([]byte(meta), &config)
	if metalen > 0 && !isYamlEmpty(config) {
		if config.Title == "" {
			title = getTitie(string(fileread))
		} else {
			title = config.Title
		}
		summary = config.Summary
		update = config.Update
		date = config.Date
		if len(config.Authors) > 0 {
			author = string(blackfriday.MarkdownCommon([]byte(strings.Join(config.Authors, " "))))
		} else {
			author = string(blackfriday.MarkdownCommon([]byte(config.Author)))
		}
		body = string(blackfriday.MarkdownCommon(fileread[metalen:]))
	}
	if metalen == 0 || isYamlEmpty(config) {
		// 按行解析元数据,元数据长度为0并且解析为空
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

// 返回data中的元数据和元数据所占的长度
func getMeta(data string) (string, int) {
	re := regexp.MustCompile(`---([\s\S]*?\n)---`)
	return re.FindString(data), len(re.FindString(data))
}

// 返回data中一级标题
func getTitie(data string) string {
	re := regexp.MustCompile(`(#)[^\n]*?\n`)
	return strings.ReplaceAll(re.FindString(data), "# ", "")
}

func isYamlEmpty(config baseconfig) bool {
	return config.Title == "" && config.Date == "" && config.Update == "" && config.Summary == "" && len(config.Author) == 0
}
