package main

import (
	"github.com/marceldegraaf/smartmeter/parser"
	"github.com/marceldegraaf/smartmeter/poller"
)

func main() {
	poller.Initialize()

	go poller.Poll()

	for {
		select {
		case msg := <-poller.Incoming:
			parser.Parse(msg)
		}
	}
}
