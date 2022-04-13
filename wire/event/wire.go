//+build wireinject

package event

import (
	"github.com/google/wire"
)

func InitializeEvent(phrase, s string) (Event, error) {
	wire.Build(NewEvent, NewGreeter, NewMessage)
	return Event{}, nil
}
