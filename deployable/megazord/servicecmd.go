package main

import (
	"github.com/state-alchemists/ayanami/config"
	"github.com/state-alchemists/ayanami/msgbroker"
	"github.com/state-alchemists/ayanami/service"
	"log"
)

// ServiceCmd declaration
func ServiceCmd() {
	serviceName := "cmd"
	// define broker
	broker, err := msgbroker.NewNats(config.GetNatsURL())
	if err != nil {
		log.Fatal(err)
	}
	// define services
	services := service.Services{ 
		service.NewCmd(serviceName, "figlet",
			[]string{"/bin/sh", "-c", "figlet $input"},
		),
	}
	// consume and publish forever
	forever := make(chan bool)
	services.ConsumeAndPublish(broker, serviceName)
	<-forever
}