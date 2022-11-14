/*
 * @Author: ning.huang ning.huang
 * @Date: 2022-11-14 09:08:27
 * @LastEditors: ning.huang ning.huang
 * @LastEditTime: 2022-11-14 09:42:02
 * @FilePath: \code\goLang\beego-examples\WebIM\routers\router.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package routers

import (
	"WebIM/controllers"

	"github.com/astaxie/beego"
)

func init() {
	// Register routers.
	beego.Router("/", &controllers.AppController{})
	// Indicate AppController.Join method to handle POST requests.
	beego.Router("/join", &controllers.AppController{}, "post:Join")

	// Long polling.
	beego.Router("/lp", &controllers.LongPollingController{}, "get:Join")
	beego.Router("/lp/post", &controllers.LongPollingController{})
	beego.Router("/lp/fetch", &controllers.LongPollingController{}, "get:Fetch")

	// WebSocket.
	beego.Router("/ws", &controllers.WebSocketController{})
	beego.Router("/ws/join", &controllers.WebSocketController{}, "get:Join")

}
