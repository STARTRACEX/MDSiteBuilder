# функция go

2022-09-30
2022-10-30
[STARTRACE](/)

Эти функции находятся в / model / docs.go

```go

func MarkdownPost(mdFilePath string) Post {}

```

 Эта функция возвращает введённую строку как адрес в аналитический целевой документ

```go

func RenderPost(c *gin.Context) {}

```

Эта функция будет объединять файлы cookie ("lang") и путь запроса для анализа файлов и отображения шаблонов HTML

```go

func RenderPostAssigned(c *gin.Context, Assigned) {}

```

 Эта функция будет объединять пути запроса с адресом для анализа документа и отображения в шаблоне HTML

```go

func RenderCatalog(DirectoryFilePath string) []Cata {}

```

 Эта функция считывает входящую DirectoryFilePath и делит символ "\\" на разделитель, делит отступ, заголовок, путь к URL, возвращает фрагмент данных
