# MDStaticSiteBuilder

1970-01-01
2022-11-01
[STARTRACEX](https://github.com/STARTRACEX/MDStaticSiteBuilder)

## Getting Started

[LIBRARY](https://github.com/STARTRACEX/MDStaticSiteBuilder)

### URL as Path

*By default*:

 The document contains the same prefix "/docs",THE ROUTE before it (such as en or zh) will be mapped to the local "./posts/THE ROUTE/documents"

"/" will return the readme file of the root directory "./README.md"

"/docs/..." will return the file with the address of path and. md as the path. If the file does not exist, the path will be spliced with index.md. If it still does not exist, the path will be spliced with "/index.md"

### Out-of-the-box

Start the service by running the executable

Set the port number and domain name by carrying the following flags

```terminal
-p <port number>
-d <domain name>

./main.exe -p 80 -d localhost
```

Basic styles and scripts have been prepared:
Including multilingual, theme, copybox directory, progress, source...
