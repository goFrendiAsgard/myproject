package main

import (
	"github.com/goFrendiAsgard/myproject/gen"
)

func init() {

	// define service
	service := gen.NewEmptyGoService(&Gen,
		"html",                                  // service name
		"github.com/goFrendiAsgard/servicehtml", // repo name
	)

	// add html.pre
	service.Set("pre", gen.NewFunction("html", "pre",
		[]string{"input"},  // inputs
		[]string{"output"}, // outputs
		[]string{},         // dependencies
	))

	// register service
	Gen.AddConfig(service)

}
