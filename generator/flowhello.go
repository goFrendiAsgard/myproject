package main

import (
	"github.com/goFrendiAsgard/myproject/gen"
)

func init() {

	// define flow
	flow := gen.NewEmptyFlow(&Gen,
		"github.com/goFrendiAsgard/flowhello", // repo name
		"hello",                               // flow name
		[]string{"content", "code"},           // inputs
		[]string{"content", "code"},           // outputs
	)

	// get http request from "/hello" and send 200 as http response's code
	flow.AppendEvent(gen.Event{
		InputEventName:  "trig.request.get.hello.out",
		UseValue:        true,
		Value:           200,
		VarName:         "code",
		OutputEventName: "trig.response.get.hello.in.code",
	})

	// get http request from "/hello", send it to  and send "Hello there" as http response's content
	flow.AppendEvent(gen.Event{
		InputEventName:  "trig.request.get.hello.out.requestURI",
		UseFunction:     true,
		FunctionPackage: "greeting",
		FunctionName:    "Hello",
		VarName:         "content",
		OutputEventName: "trig.response.get.hello.in.content",
	})

	// register flow
	Gen.AddConfig(flow)

}
