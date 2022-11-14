/*
 * @Author: ning.huang ning.huang
 * @Date: 2022-11-14 14:21:34
 * @LastEditors: ning.huang ning.huang
 * @LastEditTime: 2022-11-14 14:51:12
 * @FilePath: \code\goLang\distributed\service\service.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package service

import (
	"context"
	"fmt"
	"log"
	"net/http"
)

func Start(ctx context.Context, serviceName,
	host, port string, registerHandlersFunc func()) (context.Context, error) {

	ctx = startService(ctx, serviceName, host, port)

	registerHandlersFunc()

	return ctx, nil
}

func startService(ctx context.Context, serviceName, host, port string) context.Context {

	ctx, cancel := context.WithCancel(ctx)
	var srv http.Server
	srv.Addr = ":" + port
	go func() {
		log.Println(srv.ListenAndServe())
		cancel()
	}()

	go func() {
		fmt.Printf("%v started. \n", serviceName)
		var s string
		fmt.Scanln(&s)
		srv.Shutdown(ctx)
		cancel()
	}()

	return ctx
}
