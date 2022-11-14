package main

import (
	"context"
	"distributed/grades"
	"distributed/registry"
	"distributed/service"
	"fmt"
	stlog "log"
)

func main() {

	host, port := "localhost", "6000"
	serverAddress := fmt.Sprintf("http://%s:%s", host, port)
	r := registry.Registration{
		ServiceName: registry.GradingService,
		ServiceURL:  serverAddress,
	}
	ctx, err := service.Start(
		context.Background(),
		r,
		host,
		port,
		grades.RegisterHandlers,
	)
	if err != nil {
		stlog.Fatal("fuck you")
	}

	<-ctx.Done()

	fmt.Println("Shuting down grade service..")
}
