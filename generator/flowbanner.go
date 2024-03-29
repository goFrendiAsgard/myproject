package main

import (
	"github.com/goFrendiAsgard/myproject/gen"
)

func init() {

	// define flow
	flow := gen.NewEmptyFlow(&Gen,
		"github.com/goFrendiAsgard/flowbanner", // repo name
		"banner",                               // flow name
		[]string{"request"},                    // inputs
		[]string{"content", "code"},            // outputs
	)

	// get http request from "/banner"
	flow.AppendEvent(gen.Event{
		InputEventName: "trig.request.get.banner.out",
		VarName:        "request",
	})

	// send request.form.text.0 to figlet
	flow.AppendEvent(gen.Event{
		OutputEventName: "srvc.cmd.figlet.in.input",
		VarName:         "request.form.text.0",
	})

	// get output from figlet and send it to pre
	flow.AppendEvent(gen.Event{
		InputEventName:  "srvc.cmd.figlet.out.output",
		VarName:         "figletOut",
		OutputEventName: "srvc.html.pre.in.input",
	})

	// get output from pre and send it as http response's content
	flow.AppendEvent(gen.Event{
		InputEventName:  "srvc.html.pre.out.output",
		VarName:         "content",
		OutputEventName: "trig.response.get.banner.in.content",
	})

	// also, set http response's code into 200
	flow.AppendEvent(gen.Event{
		InputEventName:  "srvc.html.pre.out.output",
		UseValue:        true,
		Value:           200,
		VarName:         "code",
		OutputEventName: "trig.response.get.banner.in.code",
	})

	// the rest of this are error handlers

	// get error from pre and send "Internal Server Error" as http response's content
	flow.AppendEvent(gen.Event{
		InputEventName:  "srvc.html.pre.err.message",
		UseValue:        true,
		Value:           "Internal Server Error",
		VarName:         "content",
		OutputEventName: "trig.response.get.banner.in.content",
	})

	// also, set http response's code into 500
	flow.AppendEvent(gen.Event{
		InputEventName:  "srvc.html.pre.err.message",
		UseValue:        true,
		Value:           500,
		VarName:         "code",
		OutputEventName: "trig.response.get.banner.in.code",
	})

	// get error from figlet and send "Internal Server Error" as http response's content
	flow.AppendEvent(gen.Event{
		InputEventName:  "srvc.cmd.figlet.err.message",
		UseValue:        true,
		Value:           "Internal Server Error",
		VarName:         "content",
		OutputEventName: "trig.response.get.banner.in.content",
	})

	// also, set http response's code into 500
	flow.AppendEvent(gen.Event{
		InputEventName:  "srvc.cmd.figlet.err.message",
		UseValue:        true,
		Value:           500,
		VarName:         "code",
		OutputEventName: "trig.response.get.banner.in.code",
	})

	// register flow
	Gen.AddConfig(flow)
}
