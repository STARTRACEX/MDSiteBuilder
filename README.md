# MDStaticSiteBuilder

2022-11-01
2022-12-24
[STARTRACEX](https://github.com/STARTRACEX/MDStaticSiteBuilder)

## Getting Started

This is a part of [STARTRACEX.GITHUB.IO](https://startracex.github.io/)

[LIBRARY](https://github.com/STARTRACEX/MDStaticSiteBuilder)

### URL as Path

The following rules will be followed by default

- "/,/zh,/ru" will return the readme of the root directory "./README.md,./README. zh.md,./README.ru.md"
- The document will be matched with the prefix "/docs". The path before it (such as en or zh) and the path after it (the relative file path) will be mapped to the local "./posts/...", for example (./posts/ru/a/b.md will be matched by/ru/docs/a/b)
- If the file as the path does not exist after splicing, the path will be spliced with index.md. If the file still does not exist, the path will be spliced with "/index.md". If the file does not exist, all documents with default values will be returned
- The metadata of the document can be the yaml content between three dashes and the newline. If the metadata does not exist, each line of the text will be used as the metadata, or the default value will be maintained
- The catal.txt document under each language name directory will become the sidebar directory when using this language, except the content under the root directory

### Out-of-the-box

Start the service by running the executable

Set the port number and domain name by carrying the following flags

```terminal
-P<port number>
-D<domain name>
-Dev<enable the output of gin>

Example
./main.exe -p 80 -d localhost -dev
```

After running [Visit this page] (http: localhost: 8080/)

Basic styles and scripts have been prepared:
Including multilingual, theme, copybox directory, progress, source...
