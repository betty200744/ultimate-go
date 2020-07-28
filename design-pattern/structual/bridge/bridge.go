package bridge

import "fmt"

/*

桥接模式分离抽象部分和实现部分。使得两部分独立扩展。

桥接模式类似于策略模式，区别在于策略模式封装一系列算法使得算法可以互相替换。

策略模式使抽象部分和实现部分分离，可以独立变化。

*/
type AbstractMessage interface {
	SendMessage(text, to string)
}

type MessageImplementer interface {
	Send(text, to string)
}

type SMS struct {
}

func ViaSMS() MessageImplementer {
	return &SMS{}
}
func (*SMS) Send(msg, to string) {
	fmt.Printf("Send msg %v to %v via SMS \n", msg, to)
}

type Email struct {
}

func ViaEmail() MessageImplementer {
	return &Email{}
}

func (*Email) Send(msg, to string) {
	fmt.Printf("Send msg %v to %v via Email \n", msg, to)
}

type CommonMessage struct {
	via MessageImplementer
}

func NewCommentMessage(via MessageImplementer) *CommonMessage {
	return &CommonMessage{via: via}
}
func (m *CommonMessage) SendMessage(msg, to string) {
	m.via.Send(msg, to)
}

type Urgency struct {
	via MessageImplementer
}

func NewUrgencyMessage(via MessageImplementer) *Urgency {
	return &Urgency{via: via}
}
func (m *Urgency) SendMessage(msg, to string) {
	m.via.Send(msg, to)
}
