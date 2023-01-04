package model

import (
	"md/config"
	"os"
	"path"
	"strings"
)

// 侧边目录
type Cat struct {
	Title string
	Level int8
	Href  string
}

// 遍历目录
func WalkCata(ymlpath string) []Cat {
	var c = new([]Cat)
	walkCata(ymlpath, 1, c)
	return *c
}

func walkCata(ymlpath string, level int8, catList *[]Cat) *[]Cat {
	data, _ := os.ReadFile(ymlpath)
	lines := strings.Split(string(data), "\n")
	for _, line := range lines {
		line = strings.ReplaceAll(line, "\r", "")
		s := strings.Split(line, ":")
		direntry, _ := os.ReadDir(path.Dir(ymlpath))
		// 全部文件
		for _, fi := range direntry {
			if !fi.IsDir() {
				// 获取后缀为.md的文件
				if path.Ext(fi.Name()) == ".md" {
					if s[0] == strings.TrimSuffix(fi.Name(), ".md") {
						*catList = append(*catList, Cat{
							Level: level,
							Title: s[1],
							Href:  MaselHref("/" + path.Join(path.Dir(ymlpath), fi.Name())),
						})
					}
				}
			}
		}
		if IsExist(path.Join(path.Dir(ymlpath), s[0], ".yml")) {
			walkCata(path.Join(path.Dir(ymlpath), s[0], ".yml"), level+1, catList)
		}
	}
	return catList
}

func MaselHref(str string) string {
	str = strings.TrimSuffix(str, "/index.md")
	str = strings.TrimSuffix(str, ".md")
	if strings.HasPrefix(str, config.PostsPath+"/zh") {
		str = strings.Replace(str, "/posts/zh", "/zh/docs", 1)
	} else if strings.HasPrefix(str, "/posts/en") {
		str = strings.Replace(str, "/posts/en", "/en/docs", 1)
	} else {
		str = strings.Replace(str, "/posts/ru", "/ru/docs", 1)
	}
	return str
}
