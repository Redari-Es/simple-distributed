package main

import (
	"context"
	"fmt"
	stlog "log"

	"github.com/shonh/distributed/grades"
	"github.com/shonh/distributed/registry"
	"github.com/shonh/distributed/service"
)

func main() {
	host, port := "localhost", "6000"
	serviceAddress := fmt.Sprintf("http://%v:%v", host, port)

	r := registry.Registration{
		ServiceName: registry.GradingService,
		ServiceURL:  serviceAddress,
	}
	ctx, err := service.Start(context.Background(),
		host,
		port,
		r,
		grades.RegisterHandlers)
	if err != nil {
		stlog.Fatal(err)
	}
	<-ctx.Done()
	fmt.Println("Shutting down grading service")
}
