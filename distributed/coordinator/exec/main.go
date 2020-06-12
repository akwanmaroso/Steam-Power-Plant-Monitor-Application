package main

import (
	"fmt"
	"github.com/akwanmaroso/monitoring/distributed/coordinator"
)

var dc *coordinator.DatabaseConsumer
var wc *coordinator.WebappConsumer

func main() {
	ea := coordinator.NewEventAggregator()
	dc = coordinator.NewDatabaseConsumer(ea)
	wc = coordinator.NewWebappConsumer(ea)
	ql := coordinator.NewQueueListener(ea)
	go ql.ListenForNewSource()

	var a string
	fmt.Scanln(&a)
}
