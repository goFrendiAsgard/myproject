package main

func main() {
	forever := make(chan bool)
	go FlowBanner()
	go FlowHello()
	go FlowRoot()
	go GatewayGateway()
	go ServiceCmd()
	go ServiceHtml()
	<-forever
}