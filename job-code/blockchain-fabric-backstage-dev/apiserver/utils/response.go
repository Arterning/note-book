package utils

import (
	"encoding/json"
	"github.com/crc/zlzk/apiserver/define"
	"github.com/gin-gonic/gin"
	"github.com/op/go-logging"
	"net/http"
)

var Logger = logging.MustGetLogger("apiServer")

func Response(response interface{}, c *gin.Context, status int) {
	c.Writer.Header().Set("content-Type", "application/json")
	var msg []byte
	if status == http.StatusOK {
		res := define.Response{
			Code:    200,
			Message: "SUCCESS",
			Data:    response.(string),
		}
		msg, _ = json.Marshal(res)
	} else {
		res := define.Response{
			Code:    -1,
			Message: response.(string),
			Data:    "",
		}
		msg, _ = json.Marshal(res)
		Logger.Errorf(string(msg))
	}
	c.Writer.WriteHeader(status)
	c.Writer.Write(msg)
}
