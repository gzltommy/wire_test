package main

import "fmt"

type Message struct {
	msg string
}
type Greeter struct {
	Message Message
}
type Event struct {
	Greeter Greeter
}

// NewMessage Message 的构造函数
func NewMessage(msg string) Message {
	return Message{
		msg: msg,
	}
}

// NewGreeter Greeter 构造函数
func NewGreeter(m Message) Greeter {
	return Greeter{Message: m}
}

// NewEvent Event 构造函数
func NewEvent(g Greeter) Event {
	return Event{Greeter: g}
}
func (e Event) Start() {
	msg := e.Greeter.Greet()
	fmt.Println(msg)
}
func (g Greeter) Greet() Message {
	return g.Message
}

//// 使用 wire 前
//func main() {
//	message := NewMessage("hello world")
//	greeter := NewGreeter(message)
//	event := NewEvent(greeter)
//
//	event.Start()
//}

// 使用 wire 后
func main() {
	event := InitializeEvent("hello_world")

	event.Start()
}
