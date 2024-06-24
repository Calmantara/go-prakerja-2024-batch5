package order

import (
	"fmt"
	"time"
)

var (
	orderStatePending = "Pending"
	orderStatePaid    = "Paid"
	OrderMerchant     = "PREMIUM"
)

type Order struct {
	Number   int
	Merchant string
	state    string
}

func CreateOrder() Order {
	order := Order{state: orderStatePending, Merchant: OrderMerchant}
	order.generateNumber()
	fmt.Println(order.state)
	return order
}

func (o *Order) generateNumber() {
	o.Number = int(time.Now().UnixMilli())
}

func (o *Order) Paid() {
	o.state = orderStatePaid
}

func (o Order) GetState() string {
	return o.state
}
