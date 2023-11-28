// package message implements a simple definition about messages' format
// for system communication
package message

type MessageQueue interface {
}

type Message interface {
	Send(MessageQueue)
}
