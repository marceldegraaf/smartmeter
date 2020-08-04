package main

import (
	"marceldegraaf.net/smartmeter/parser"
	"marceldegraaf.net/smartmeter/poller"
	"marceldegraaf.net/smartmeter/store"
)

func main() {
	store.Initialize()
	poller.Initialize()

	// Poll for usage stats in a goroutine.
	go poller.Poll()

	// Any incoming channel data is sent along
	// for further processing immediately.
	for {
		select {
		case msg := <-poller.Incoming:
			parser.Parse(msg)
		case usage := <-parser.Incoming:
			store.Save(usage)
		}
	}
}
