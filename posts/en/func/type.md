---
title: Types
date: 2023-01-04
update: 2023-01-04
---

## Catalog

```go
type Cat struct {
 Title string
 Level int8
 Href  string
}
type Cata struct {
 Name  string
 Level string
 Super string
}
```

`Cat`和`Cata`是用于渲染目录的结构体,二者具有相同含义

## Post

```go
type Post struct {
 Title      string
 Date       string
 Update     string
 Summary    string
 Author     template.HTML
 Body       template.HTML
 OriginFile string // Origin File Path
}
```

`Post`是用于渲染文章的结构体,部分内容来自`Meta`

## Meta

```go
type Meta struct {
 Title   string   `yaml:"title,omitempty"`
 Author  string   `yaml:"author,omitempty"`
 Authors []string `yaml:"authors,omitempty"`
 Summary string   `yaml:"summary,omitempty"`
 Date    string   `yaml:"date,omitempty"`
 Update  string   `yaml:"update,omitempty"`
}
```

`Meta`是文档的元数据
