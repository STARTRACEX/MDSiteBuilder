---
author: "[STARTRACEX](https://github.com/STARTRACEX/MDStaticSiteBuilder)"
date: 2022-11-01
update: 2022-12-25
summary: A simple site builder
---

# MDStaticSiteBuilder

## 入门

这是[STARTRACEX.GITHUB.IO](https://startracex.github.io/)的一部分

[库](https://github.com/STARTRACEX/MDStaticSiteBuilder)

### URL作为路径

默认情况下将遵循以下规则

- “/,/zh,/ru”将返回根目录“./README.md,./README.zh.md,./README.ru.md”的自述文件
- 该文档将匹配包含前缀“/docs”,它之前的路径（如en或zh）和之后的路径（相对的文件路径）将映射到本地“./posts/...”，例如（./posts/ru/a/b.md将被/ru/docs/a/b匹配）
- 如果拼接后作为路径的文件不存在，路径将与index.md拼接。如果文件仍然不存在，则路径将与“/index.md”拼接，如果文件不存在，将返回所有文档变量均为默认值的文档
- 文档的元数据可由一三个短横线和换行之间的yaml内容作为元数据，如果元数据不存在，将尝试以文本的每行作为元数据，或是保持默认值
- 每个语言名目录下的cata.txt文档将成为使用此语言时的侧边栏目录，根目录下的内容除外
- 不以“/public”开头的资源路径将会被认为是相对与文档文件的路径

### 使用yaml

你可以在文档中使用yaml，或是使用以行分割的内容作为文档元数据，正如此文档第一行所示。

yaml解析以下内容:"author,summary,date,update,title"(字符串类型),"author"（字符串数组类型）.

**注意** author会被存在的authors覆盖，不存在的title将会是文档的一级标题，author可以使用markdown语法，请以(")包含

### 开箱即用

通过运行可执行文件启动服务
通过携带以下标志设置端口号和域名

```terminal
-p <端口号>
-d <域名>
-dev <启用gin的输出>

示例
main\main.exe -p 8080 -d localhost -dev
```

在运行后[访问此页面](http:localhost:8080/)

已准备好基本样式和脚本：
包括多语言、主题、复制框、目录、进度、源代码……
