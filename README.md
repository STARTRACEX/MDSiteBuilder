README (TITLE)
A Readme file(SUMMARY)
1970-1-1 (CREAT DATE)
2875-1-1 (LAST UPDATE)
Me (AUTHER)

## Here is a Markdown Site Easy Quick Builder

[Go to Watch](https://github.com/STARTRACEX/MDStaticSiteBuilder)

## Rules

By *default*, "/" will match the readme of the root directory The first parameter carried by md and "/" will be the default language and may not exist. The subsequent path will match the file path.

1. Create a folder named language in./posts as the matching directory for this language.

2. If the "default" folder exists, it will be the default language matching directory; otherwise, the default language matching directory will be the folder in./posts (excluding the folder name of the declared language).

3. If the path points to a folder but does not point to a file, the index.md in this folder will be searched.

4. If the path points to a file or folder with the same name, the file will be matched.

## Identify document information

In the rendered document, it is divided into two parts.

Part I: the first line: the title of the document, the second line: the summary of the document, the third line: the creation time of the document, the fourth line: the last modification time of the document, the fifth line: the author of the document.

The following content is the second part: the main body of the document (please ensure that the first part of each document exists).

## **Watch Out**

- **All documents must have lines exceed 4**

- **All document supports maximum directory level 2**
