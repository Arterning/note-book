/**

go mod init shorturl

go run main.go
提示你需要go get...

go get github.com/astaxie/beego
go get github.com/beego/samples/shorturl/controllers

go mod download

**/

/**
go mod 命令
1. 拉取缺少的模块，移除不用的模块 ： go mod tidy
2. 下载依赖 : go mod download
3. 检验依赖: go mod verify
4. 查看命令列表: go mod
**/