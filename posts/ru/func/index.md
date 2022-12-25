# функция go

2022-09-30
2022-11-01
[STARTRACE](/)

Эти функции находятся в ./model

```go

func RenderMarkdown(mdFilePath string) Post {}

```

 Эта функция возвращает введённую строку как адрес в аналитический целевой документ

```go

func RenderPost(c *gin.Context) {}

```

Эта функция будет объединять файлы "lang" и путь запроса для анализа файлов и отображения шаблонов HTML

```go

func RenderCatalog(DirectoryFilePath string) []Cata {}

```

 Эта функция считывает входящую DirectoryFilePath и делит символ "\\" на разделитель, делит отступ, заголовок, путь к URL, возвращает фрагмент данных
