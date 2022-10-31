# Functions for go

2022-09-30
2022-10-30
[STARTRACE](/)

These functions can be found in./model/docs.go

```go
func MarkdownPost(mdFilePath string) Post {}
```

This function returns the parsed target document with the passed in string as the address

```go
func RenderPost(c *gin.Context) {}
```

This function will combine cookies ("lang") and request paths to resolve documents for addresses, and render them to HTML templates

```go
func RenderPostAssigned(c *gin.Context, Assigned) {}
```

This function will combine Assigned and request path to parse documents for addresses and render them to HTML templates

```go
func RenderCatalog(DirectoryFilePath string) []Cata {}
```

This function will read the incoming DirectoryFilePath, use " " as the dividing symbol, split the indent level, title, URL path, and return the split data slice
