# MDStaticSiteBuilder

1970-01-01
2022-11-01
[STARTRACEX](https://github.com/STARTRACEX/MDStaticSiteBuilder)

## 入门

[库]<https://github.com/STARTRACEX/MDStaticSiteBuilder>)

### URL作为路径

*默认情况下*
该文档包含相同的前缀“/docs”,它之前的路由（如en或zh）将映射到本地“./posts/THE ROUTE/documents”
“/”将返回根目录“./readme.md”的自述文件
“/docs/…”将路径与“.md拼接”，如果拼接后作为路径的文件不存在，路径将与index.md拼接。如果文件仍然不存在，则路径将与“/index.md”拼接，最终返回文件

### 开箱即用

通过运行可执行文件启动服务
通过携带以下标志设置端口号和域名

```terminal
-p<端口号>
-d<域名>

./main.exe -p 80 -d localhost
```

已准备好基本样式和脚本：
包括多语言、主题、复制框、目录、进度、源代码……
