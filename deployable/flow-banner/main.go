package main


import (
	"github.com/state-alchemists/ayanami/config"
	"github.com/state-alchemists/ayanami/msgbroker"
	"github.com/state-alchemists/ayanami/service"
	"log"
)

// main declaration
func main() {
	// define broker
	broker, err := msgbroker.NewNats(config.GetNatsURL())
	if err != nil {
		log.Fatal(err)
	}
	// define services
	services := service.Services{
		service.NewFlow("banner", broker,
			// inputs
			[]string{ "request" },
			// outputs
			[]string{ "content", "code" },
			[]service.FlowEvent{ 
				{ 
					InputEvent: "trig.request.get.banner.out",
					VarName:    "request",
				},
				{ 
					OutputEvent: "srvc.cmd.figlet.in.input",
					VarName:     "request.form.text.0",
				},
				{ 
					InputEvent:  "srvc.cmd.figlet.out.output",
					OutputEvent: "srvc.html.pre.in.input",
					VarName:     "figletOut",
				},
				{ 
					InputEvent:  "srvc.html.pre.out.output",
					OutputEvent: "trig.response.get.banner.in.content",
					VarName:     "content",
				},
				{ 
					InputEvent:  "srvc.html.pre.out.output",
					OutputEvent: "trig.response.get.banner.in.code",
					UseValue:    true,
					Value:       200,
					VarName:     "code",
				},
				{ 
					InputEvent:  "srvc.html.pre.err.message",
					OutputEvent: "trig.response.get.banner.in.content",
					UseValue:    true,
					Value:       "Internal Server Error",
					VarName:     "content",
				},
				{ 
					InputEvent:  "srvc.html.pre.err.message",
					OutputEvent: "trig.response.get.banner.in.code",
					UseValue:    true,
					Value:       500,
					VarName:     "code",
				},
				{ 
					InputEvent:  "srvc.cmd.figlet.err.message",
					OutputEvent: "trig.response.get.banner.in.content",
					UseValue:    true,
					Value:       "Internal Server Error",
					VarName:     "content",
				},
				{ 
					InputEvent:  "srvc.cmd.figlet.err.message",
					OutputEvent: "trig.response.get.banner.in.code",
					UseValue:    true,
					Value:       500,
					VarName:     "code",
				},
			},
		),
	}
	// consume and publish forever
	forever := make(chan bool)
	services.ConsumeAndPublish(broker, "flow")
	<-forever
}