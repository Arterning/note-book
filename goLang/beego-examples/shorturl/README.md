<!--
 * @Author: error: git config user.name && git config user.email & please set dead value or install git
 * @Date: 2022-11-14 09:08:27
 * @LastEditors: error: git config user.name && git config user.email & please set dead value or install git
 * @LastEditTime: 2022-11-14 09:22:50
 * @FilePath: \code\goLang\beego-examples\shorturl\README.md
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
-->
# Shorturl

[中文文档](./README_ZH.md)

This sample is a API application based on beego. It has two API func:

- /v1/shorten
- /v1/expand

## Installation

```
cd $GOPATH/src/samples/shorturl
bee run
```

## Usage:

```
# shortening url example
http://localhost:8080/v1/shorten/?longurl=http://google.com

{
  "UrlShort": "5laZG",
  "UrlLong": "http://google.com"
}

# expanding url example
http://localhost:8080/v1/expand/?shorturl=5laZG

{
  "UrlShort": "5laZG",
  "UrlLong": "http://google.com"
}
```
