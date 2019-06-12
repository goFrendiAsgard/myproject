package main

import (
	"github.com/state-alchemists/ayanami/config"
	"github.com/state-alchemists/ayanami/gateway"
	"github.com/state-alchemists/ayanami/msgbroker"
	"log"
)

// main declaration
func main() {
	routes := []string{ // define your routes here
		"/",
		"/hello/",
		"/banner/",
	}
	broker, err := msgbroker.NewNats(config.GetNatsURL())
	if err != nil {
		log.Fatal(err)
	}
	port := config.GetGatewayPort()
	multipartFormLimit := config.GetGatewayMultipartFormLimit()
	gateway.Serve(broker, port, multipartFormLimit, routes)
}