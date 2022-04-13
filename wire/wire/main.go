package main

import (
	"go-tips/wire/event"
)

func main() {
	e, err := event.InitializeEvent("hello phrase")
	if err != nil {
		panic(err)
	}
	e.Start()
}
