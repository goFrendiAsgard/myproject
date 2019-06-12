package main

import (
	"github.com/goFrendiAsgard/myproject/gen"
)

func init() {

	// define procedure
	procedure := gen.NewGoMonolith(&Gen,
		"megazord",                 // app directory
		"github.com/nerv/megazord", // repo name
	)

	// register procedure
	Gen.AddProcedure(procedure)

}