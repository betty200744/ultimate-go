package strategy

import "fmt"

/*
 支付方式策略
*/

type Payment struct {
	context  *PaymentContext
	strategy PaymentStrategy
}

func NewPayment(name, cardId string, money int, strategy PaymentStrategy) *Payment {
	return &Payment{
		context: &PaymentContext{
			name:   name,
			cardId: cardId,
			money:  money,
		},
		strategy: strategy,
	}
}
func (p *Payment) Pay() {
	p.strategy.Pay(p.context)
}

type PaymentContext struct {
	name   string
	cardId string
	money  int
}

type PaymentStrategy interface {
	Pay(*PaymentContext)
}

type Cash struct {
}

func (*Cash) Pay(ctx *PaymentContext) {
	fmt.Printf("Cash Pay $%d to %s by cash \n", ctx.money, ctx.name)
}

type Bank struct {
}

func (*Bank) Pay(ctx *PaymentContext) {
	fmt.Printf("Bank Pay $%d to %s by cash \n", ctx.money, ctx.name)
}
