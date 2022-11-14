<!--
 * @Author: ning.huang ning.huang
 * @Date: 2022-11-14 09:08:27
 * @LastEditors: ning.huang ning.huang
 * @LastEditTime: 2022-11-14 10:11:07
 * @FilePath: \code\goLang\beego-examples\todo\README.md
-->
# todo list

angularJS + beego demo

## Installation

```
cd $GOPATH/src/samples/todo
bee run
```

## Usage

```
http://127.0.0.1:8080

curl -H "Content-Type: application/json" -X POST -d '{"Title": "Buy bread"}' \
"http://127.0.0.1:8080/task"
 
curl -X GET "http://127.0.0.1:8080/task"

```

