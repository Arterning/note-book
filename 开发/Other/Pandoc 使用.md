Pandoc是一款非常强大的文档格式转换工具，对于Word转Markdown的场景，可以快速的实现转换，并且把Word文件中图片，生成到指定的文件夹中。

> Pandoc是由John MacFarlane开发的标记语言转换工具，可实现不同标记语言间的格式转换，堪称该领域中的“瑞士军刀”。

Pandoc understands a number of useful markdown syntax extensions, including document metadata (title, author, date); footnotes; tables; definition lists; superscript and subscript; strikeout; enhanced ordered lists (start number and numbering style are significant); running example lists; delimited code blocks with syntax highlighting; smart quotes, dashes, and ellipses; markdown inside HTML blocks; and inline LaTeX. If strict markdown compatibility is desired, all of these extensions can be turned off.

上面是引用的，关于Pandoc的介绍。

```javascript
pandoc -f docx -t markdown test.docx -o test.md --extract-media ./images
```


Pandoc还支持，直接通过Http协议访问网页的内容，并生成Markdown文件，远端Web[服务器](https://cloud.tencent.com/act/pro/promotion-cvm?from_column=20065&from=20065)上的图片也可以，保存到本地的图片文件夹中。

```javascript
pandoc -f html -t markdown --request-header User-Agent:“Mozilla/5.0” https://candylab.net/design/HFishSOC/ -o candylab.md --extract-media ./images1
```

