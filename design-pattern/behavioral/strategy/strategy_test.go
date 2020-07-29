package strategy

import "testing"

func TestBank_Pay(t *testing.T) {
	strategy := Bank{}
	pay := NewPayment("betty", "111", 100, &strategy)
	pay.strategy.Pay(pay.context)
}
func TestCash_Pay(t *testing.T) {
	strategy := Cash{}
	pay := NewPayment("betty", "111", 100, &strategy)
	pay.strategy.Pay(pay.context)
}
