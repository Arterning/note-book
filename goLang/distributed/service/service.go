/*
 * @Author: ning.huang ning.huang
 * @Date: 2022-11-14 14:21:34
 * @LastEditors: ning.huang ning.huang
 * @LastEditTime: 2022-11-14 17:12:19
 * @FilePath: \code\goLang\distributed\service\service.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package service

import (
	"context"
	"distributed/registry"
	"fmt"
	"log"
	"net/http"
)

// Start用于启动一个服务
// registerHandlersFunc是服务的逻辑
func Start(ctx context.Context, reg registry.Registration,
	host, port string, registerHandlersFunc func()) (context.Context, error) {

	registerHandlersFunc()
	ctx = startService(ctx, reg.ServiceName, host, port) //启动服务
	err := registry.RegisterService(reg)                 //注册服务
	if err != nil {
		return ctx, err
	}

	return ctx, nil
}

func startService(ctx context.Context, serviceName registry.ServiceName, host, port string) context.Context {

	ctx, cancel := context.WithCancel(ctx)
	var srv http.Server
	srv.Addr = ":" + port
	go func() {
		log.Println(srv.ListenAndServe())
		err := registry.ShutdownService(fmt.Sprintf("http://%s:%s", host, port))
		if err != nil {
			log.Println(err)
		}
		cancel()
	}()

	go func() {
		fmt.Printf("%v started. \n", serviceName)
		var s string
		fmt.Scanln(&s)
		err := registry.ShutdownService(fmt.Sprintf("http://%s:%s", host, port))
		if err != nil {
			log.Println(err)
		}
		srv.Shutdown(ctx)
		cancel()
	}()

	return ctx
}
