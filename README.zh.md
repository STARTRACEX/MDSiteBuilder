# MDStaticSiteBuilder

2022-11-01
2022-12-24
[STARTRACEX](https://github.com/STARTRACEX/MDStaticSiteBuilder)

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

### 开箱即用

通过运行可执行文件启动服务
通过携带以下标志设置端口号和域名

```terminal
-p <端口号>
-d <域名>
-dev <启用gin的输出>

示例
./main.exe -p 80 -d localhost -dev
```

在运行后[访问此页面](http:localhost:8080/)

已准备好基本样式和脚本：
包括多语言、主题、复制框、目录、进度、源代码……
