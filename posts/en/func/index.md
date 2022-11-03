# Functions for go

2022-09-30
2022-11-01
[STARTRACE](/)

These functions can be found in./model/docs.go
![这是图片](../img/1.jpg "Magic Gardens")

```go
func RenderMarkdown(mdFilePath string) Post {}
```

This function returns the parsed target document with the passed in string as the address

```go
func RenderPost(c *gin.Context) {}
```

This function will combine "lang" and request paths to resolve documents for addresses, and render them to HTML templates

```go
func RenderCatalog(DirectoryFilePath string) []Cata {}
```

This function will read the incoming DirectoryFilePath, use " " as the dividing symbol, split the indent level, title, URL path, and return the split data slice
