package main

import (
	"github.com/goFrendiAsgard/myproject/gen"
)

func init() {

	// define flow
	flow := gen.NewEmptyFlow(&Gen,
		"github.com/goFrendiAsgard/flowroot", // repo name
		"root",                               // flow name
		[]string{"content", "code"},          // inputs
		[]string{"content", "code"},          // outputs
	)

	// get http request from "/" and send 200 as http response's code
	flow.AppendEvent(gen.Event{
		InputEventName:  "trig.request.get.out",
		UseValue:        true,
		Value:           200,
		VarName:         "code",
		OutputEventName: "trig.response.get.in.code",
	})

	// get http request from "/" and send "Hello there" as http response's content
	flow.AppendEvent(gen.Event{
		InputEventName:  "trig.request.get.out",
		UseValue:        true,
		Value:           "Hello there",
		VarName:         "content",
		OutputEventName: "trig.response.get.in.content",
	})

	// register flow
	Gen.AddConfig(flow)

}
