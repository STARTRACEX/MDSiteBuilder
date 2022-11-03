# Use markdown in go

2022-09-30
2022-10-30
[STARTRACE](/)

## Getting Started

![这是图片](img/1.jpg "Magic Gardens")

### Use directory

Write directory file (saved in posts/zh/data.txt)
The format is &lt;indent level&gt;\\&lt;title&gt;\\&lt;URL path&gt;\\
Example

```txt

1\TITLE\/en/docs/urlpath

```

### Multiterminal rendering

**Server rendering**\

The server will render the document and directory structure converted to markup language

**Client Rendering**\

The client will render the title directory of a single document, which is executed by javascript

***Code highlighting and theme colors***

Third party javascript plug-in (Highlight. js) and style parser (Less. js) are used

### Document Rules

|Property | Line | Type|
| :---:     | :----: |:---:|
|Window title | 1 | string, remove "# "|
|Center title | 1 | string, remove "# "|
|Description/Summary | 2 | Original String|
|Creation date | 3 | original string|
|Update Date | 4 | Original String|
|Author | 5 | Markup Language|
|Document |1∩[6,∞)| Markup Language|

## Existing problems

1. If you use markdown strict mode, a warning will appear in the summary in the second line (to solve this problem, you need to leave it blank or modify the parsing method of the summary)
2. You cannot use advanced functions, such as Latex formula
