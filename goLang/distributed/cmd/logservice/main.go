package main

import (
	"context"
	"distributed/log"
	"distributed/registry"
	"distributed/service"
	"fmt"
	stlog "log"
)

func main() {

	log.Run("dis.log")

	host, port := "localhost", "4000"
	serverAddress := fmt.Sprintf("http://%s:%s", host, port)
	r := registry.Registration{
		ServiceName: "Log Service",
		ServiceURL:  serverAddress,
	}
	ctx, err := service.Start(
		context.Background(),
		r,
		host,
		port,
		log.RegisterHandlers,
	)
	if err != nil {
		stlog.Fatal("fuck you")
	}

	<-ctx.Done()

	fmt.Println("Shuting down log service..")
}
