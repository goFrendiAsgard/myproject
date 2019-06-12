package main

import (
	"github.com/goFrendiAsgard/myproject/gen"
)

func init() {

	// define service
	service := gen.NewEmptyCmd(&Gen,
		"cmd",                                  // service name
		"github.com/goFrendiAsgard/servicecmd", // repo name
	)

	// add command
	service.Set("figlet", "figlet $input")

	// register service
	Gen.AddConfig(service)

}
