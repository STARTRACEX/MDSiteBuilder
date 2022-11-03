# 用于go的功能

2022-09-30
2022-11-01
[STARTRACE](/)

这些功能位于./model/docs.go

```go
func RenderMarkdown(mdFilePath string) Post {}
```

此功能将传入的字符串作为地址返回解析的目标文档

```go
func RenderPost(c *gin.Context) {}
```

此功能将会合并"lang"与请求路径为地址解析文档，并渲染至HTML模板

```go
func RenderCatalog(DirectoryFilePath string) []Cata {}
```

此功能将读取传入的DirectoryFilePath并以"\\"为分割符号,分割缩进等级，标题，URL路径，返回分割完毕的数据切片
