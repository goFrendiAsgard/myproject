package main


import (
	"github.com/state-alchemists/ayanami/config"
	"github.com/state-alchemists/ayanami/msgbroker"
	"github.com/state-alchemists/ayanami/service"
	"log"
	"github.com/goFrendiAsgard/servicehtml/html"
)

// main declaration
func main() {
	serviceName := "html"
	// define broker
	broker, err := msgbroker.NewNats(config.GetNatsURL())
	if err != nil {
		log.Fatal(err)
	}
	// define services
	services := service.Services{ 
		service.NewService(serviceName, "pre",
			[]string{ "input" },
			[]string{ "output" },
			html.Wrappedpre,
		),
	}
	// consume and publish forever
	forever := make(chan bool)
	services.ConsumeAndPublish(broker, serviceName)
	<-forever
}