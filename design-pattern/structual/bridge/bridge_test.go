package bridge

import "testing"

func Test_NewCommentMessage(t *testing.T) {
	via := ViaSMS()
	via2 := ViaEmail()
	m := NewCommentMessage(via)
	m2 := NewCommentMessage(via2)
	m.via.Send("'hello'", "betty")
	m2.via.Send("'hello'", "betty")
}
func Test_NewUrgencyMessage(t *testing.T) {
	via := ViaSMS()
	via2 := ViaEmail()
	m := NewUrgencyMessage(via)
	m2 := NewUrgencyMessage(via2)
	m.via.Send("'hello'", "betty")
	m2.via.Send("'hello'", "betty")
}
