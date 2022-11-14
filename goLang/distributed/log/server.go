/*
 * @Author: ning.huang ning.huang
 * @Date: 2022-11-14 14:08:59
 * @LastEditors: ning.huang ning.huang
 * @LastEditTime: 2022-11-14 14:20:19
 * @FilePath: \code\goLang\distributed\log\server.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package log

import (
	"io/ioutil"
	stlog "log"
	"net/http"
	"os"
)

var log *stlog.Logger

// fileLog实际上就是string的一个别名
type fileLog string

// 给fileLog增加Write方法 这就实现了Writer接口
func (fl fileLog) Write(data []byte) (int, error) {

	f, err := os.OpenFile(string(fl), os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0600)
	if err != nil {
		return 0, err
	}
	defer f.Close()
	return f.Write(data)
}

func Run(destination string) {
	//go作为prefix
	log = stlog.New(fileLog(destination), "go", stlog.LstdFlags)
}

func RegisterHandlers() {
	http.HandleFunc("/log", func(response http.ResponseWriter, request *http.Request) {

		switch request.Method {
		case http.MethodPost:
			msg, err := ioutil.ReadAll(request.Body)
			if err != nil || len(msg) == 0 {
				response.WriteHeader(http.StatusBadRequest)
				return
			}
			write(string(msg))
		default:
			response.WriteHeader(http.StatusMethodNotAllowed)
		}
	})
}

func write(message string) {

	log.Printf("%v\n", message)

}
