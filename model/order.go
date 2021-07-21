package model

type Order struct {
	OrderID     string
	CreateTime  string
	TotalCount  int64
	TotalAmount float64
	State       int //0-未发货,1-已发货,2-交易完成，-1交易取消
	UserID      int
	OrderItems  []*OrderItem
}

func (order *Order) NoSend() bool {
	return order.State == 0
}
func (order *Order) SendComplate() bool {
	return order.State == 1
}
func (order *Order) Complate() bool {
	return order.State == 2
}
func (order *Order) Cancel() bool {
	return order.State == -1
}
