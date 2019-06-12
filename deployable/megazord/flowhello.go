package main


import (
	"github.com/state-alchemists/ayanami/config"
	"github.com/state-alchemists/ayanami/msgbroker"
	"github.com/state-alchemists/ayanami/service"
	"log"
	"github.com/nerv/megazord/greeting"
)

// FlowHello declaration
func FlowHello() {
	// define broker
	broker, err := msgbroker.NewNats(config.GetNatsURL())
	if err != nil {
		log.Fatal(err)
	}
	// define services
	services := service.Services{
		service.NewFlow("hello", broker,
			// inputs
			[]string{ "content", "code" },
			// outputs
			[]string{ "content", "code" },
			[]service.FlowEvent{ 
				{ 
					InputEvent:  "trig.request.get.hello.out",
					OutputEvent: "trig.response.get.hello.in.code",
					UseValue:    true,
					Value:       200,
					VarName:     "code",
				},
				{ 
					Function:    greeting.Hello,
					InputEvent:  "trig.request.get.hello.out.requestURI",
					OutputEvent: "trig.response.get.hello.in.content",
					UseFunction: true,
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