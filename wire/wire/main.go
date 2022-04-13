package main

import (
	"tips-go/wire/event"
)

func main() {
	e, err := event.InitializeEvent("hello phrase")
	if err != nil {
		panic(err)
	}
	e.Start()
}
