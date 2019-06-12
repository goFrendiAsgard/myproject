package main


import (
	"github.com/state-alchemists/ayanami/config"
	"github.com/state-alchemists/ayanami/msgbroker"
	"github.com/state-alchemists/ayanami/service"
	"log"
)

// FlowRoot declaration
func FlowRoot() {
	// define broker
	broker, err := msgbroker.NewNats(config.GetNatsURL())
	if err != nil {
		log.Fatal(err)
	}
	// define services
	services := service.Services{
		service.NewFlow("root", broker,
			// inputs
			[]string{ "content", "code" },
			// outputs
			[]string{ "content", "code" },
			[]service.FlowEvent{ 
				{ 
					InputEvent:  "trig.request.get.out",
					OutputEvent: "trig.response.get.in.code",
					UseValue:    true,
					Value:       200,
					VarName:     "code",
				},
				{ 
					InputEvent:  "trig.request.get.out",
					OutputEvent: "trig.response.get.in.content",
					UseValue:    true,
					Value:       "Hello there",
					VarName:     "content",
				},
			},
		),
	}
	// consume and publish forever
	forever := make(chan bool)
	services.ConsumeAndPublish(broker, "flow")
	<-forever
}