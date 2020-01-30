package main

import (
	"fmt"
)

// --------------------------------------------------------
// [Event]
//     +-- [Greeter]
//             +-- [Message]
// --------------------------------------------------------

//== Message

// Message ..
type Message string

// NewMessage ...
func NewMessage() Message {
	return Message("hello wire")
}

//== Greeter

// Greeter ...
type Greeter struct {
	Message Message
}

// NewGreeter ...
func NewGreeter(m Message) Greeter {
	return Greeter{Message: m}
}

// Greet ...
func (g Greeter) Greet() Message {
	return g.Message
}

//== Event

// Event ...
type Event struct {
	Greeter Greeter
}

// NewEvent ...
func NewEvent(g Greeter) Event {
	return Event{Greeter: g}
}

// Start ...
func (e Event) Start() {
	msg := e.Greeter.Greet()
	fmt.Println(msg)
}

// main
/*
func main() {
	message := NewMessage()
	greeter := NewGreeter(message)
	event := NewEvent(greeter)
	event.Start()
}
*/

func main() {
	e := InitializeEvent()
	e.Start()
}
