/*
 * @Author: ning.huang ning.huang
 * @Date: 2022-11-14 15:22:36
 * @LastEditors: ning.huang ning.huang
 * @LastEditTime: 2022-11-14 17:09:18
 * @FilePath: \code\goLang\distributed\cmd\regisryservice\main.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package main

import (
	"context"
	"distributed/registry"
	"fmt"
	"log"
	"net/http"
)

/*
*

	可运行的注册服务

*
*/
func main() {

	http.Handle("/services", &registry.RegistryService{})

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var srv http.Server
	srv.Addr = registry.ServerPort

	go func() {
		log.Println(srv.ListenAndServe())
		cancel()
	}()

	go func() {
		fmt.Println("Registry service started")
		var s string
		fmt.Scanln(&s)
		srv.Shutdown(ctx)
		cancel()
	}()

	<-ctx.Done()
	fmt.Println("shuting down registry service")
}
