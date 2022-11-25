package OCP

type OrderType int

const (
	APIS OrderType = iota
	OPS
)

type OrderI interface {
	Amount() int
}

type Order struct {
	t      OrderType
	amount int
}

func NewOrder(t OrderType, amount int) OrderI {
	order := &Order{
		t:      t,
		amount: amount,
	}
	switch t {
	case APIS:
		return &ApisOrder{order}
	case OPS:
		return &OpsOrder{order}
	default:
		return &ApisOrder{order}
	}
}
func (o *Order) Amount() int {
	return o.amount
}

type ApisOrder struct {
	*Order
}

type OpsOrder struct {
	*Order
}

func (o *OpsOrder) Cancel() bool {
	// do something
	return true
}

func OrderHandler(o OrderI) {

}
