/*
 * @Author: ning.huang ning.huang
 * @Date: 2022-11-14 14:56:32
 * @LastEditors: ning.huang ning.huang
 * @LastEditTime: 2022-11-14 15:33:37
 * @FilePath: \code\goLang\distributed\registry\registration.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package registry

type Registration struct {
	ServiceName ServiceName
	ServiceURL  string
}

type ServiceName string

const (
	LogService     = ServiceName("LogService")
	GradingService = ServiceName("GradingService")
)
