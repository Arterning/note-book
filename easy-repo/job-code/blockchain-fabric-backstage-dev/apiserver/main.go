package main

import (
	"flag"
	"fmt"
	"github.com/crc/gohfc"
	"github.com/crc/zlzk/apiserver/router"
	"github.com/crc/zlzk/apiserver/sdk"
	"github.com/fvbock/endless"
	"github.com/gorilla/handlers"
	"github.com/op/go-logging"
	"github.com/spf13/viper"
	"log"
	"runtime"
	"strings"
	"syscall"
)

var (
	configPath = flag.String("configPath", "", "config file path")
	configName = flag.String("configName", "", "config file name")
)

var logger = logging.MustGetLogger("apiServer")

func main() {
	// parse init param
	logger.Debug("Usage : ./apiServer -configPath= -configName=")
	flag.Parse()
	if *configPath == "" || *configName == "" {
		*configPath = "./"
		*configName = "client_sdk"
		logger.Debug("because configPath or configName nil  so  auto set")
		logger.Debug("auto set  configPath = \"./\" , configName = \"client_sdk\"")
	}

	//init sdk
	err := sdk.InitSDKs(*configPath, *configName)
	if err != nil {
		logger.Errorf("init sdk error : %s\n", err.Error())
		return
	}
	if err := sdk.SetLogLevel(viper.GetString("log.logLevel"), "apiServer"); err != nil {
		logger.Errorf("SetLogLevel error : %s\n", err.Error())
		return
	}
	go gohfc.GetHandler().CheckBlockSyncState()
	// 设置使用系统最大CPU
	runtime.GOMAXPROCS(runtime.NumCPU())
	// 构造路由器
	engine := router.GetRouter()

	// 跨域
	corsOpt := []handlers.CORSOption{
		handlers.AllowedOrigins([]string{"*"}),
		handlers.AllowedHeaders([]string{"Authorization", "Access-Control-Allow-Origin", "Origin", "X-Requested-With", "Content-Type", "Accept", "user", "password", "new_password", "token", "user_address"}),
		handlers.ExposedHeaders([]string{"Content-Length", "Access-Control-Allow-Origin"}),
		handlers.AllowedMethods([]string{"POST", "GET", "OPTIONS", "PUT", "DELETE", "PATCH"}),
		handlers.AllowCredentials(),
	}
	//Get the listen port for apiServer
	listenPort := viper.GetInt("apiServer.listenPort")
	if listenPort == 0 {
		listenPort = 8899
	}
	logger.Debug("The listen port is", listenPort)
	listenPortString := fmt.Sprintf(":%d", listenPort)

	server := endless.NewServer(listenPortString, handlers.CORS(corsOpt...)(engine))
	server.BeforeBegin = func(add string) {
		log.Printf("Actual pid is %d", syscall.Getpid())
	}
	err = server.ListenAndServe()
	if err != nil {
		if strings.Contains(err.Error(), "use of closed network connection") {
			logger.Errorf("%v\n", err)
		} else {
			logger.Errorf("Api server start failed! err:%v\n", err)
			panic(err)
		}
	}
}
