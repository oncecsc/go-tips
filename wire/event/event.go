package event

import (
	"errors"
	"fmt"
	"time"
)

type Message string

func NewMessage(phrase string) Message {
	return Message(phrase)
}

type Greeter struct {
	Grumpy  bool
	Message Message
}

func NewGreeter(m Message) Greeter {
	// fmt.Println("new greeter msg: ", s)
	grumpy := false
	if time.Now().Unix()%2 == 0 {
		grumpy = true
	}
	return Greeter{Message: m, Grumpy: grumpy}
}

func (g Greeter) Greet() Message {
	if g.Grumpy {
		return Message("Go away!")
	}
	return g.Message
}

type Event struct {
	Greeter Greeter
}

func NewEvent(g Greeter) (Event, error) {
	if g.Grumpy {
		return Event{}, errors.New("could not create event: event greeter is grumpy")
	}
	return Event{Greeter: g}, nil
}

func (e Event) Start() {
	msg := e.Greeter.Greet()
	fmt.Println(msg)
}
