package main

import (
	"github.com/goFrendiAsgard/myproject/gen"
)

func init() {

	// define gateway
	gateway := gen.NewEmptyGateway(&Gen,
		"gateway",                 // service name
		"github.com/nerv/gateway", // repo name
	)

	// add routes to gateway
	gateway.AddRoute("/")
	gateway.AddRoute("/hello/")
	gateway.AddRoute("/banner/")

	// register gateway
	Gen.AddConfig(gateway)

}
