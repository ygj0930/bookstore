package model

type Cart struct {
	ID          string
	CartItems   []*CartItem
	TotalCount  int64
	TotalAmount float64
	UserID      int
	SessionID   string
}

//获取购物车总数量
func (cart *Cart) GetTotalCount() int64 {
	totalCount := int64(0)

	for _, v := range cart.CartItems {
		totalCount = totalCount + v.Count
	}
	cart.TotalCount = totalCount
	return totalCount
}

//获取购物车总金额
func (cart *Cart) GetTotalAmount() float64 {
	totalAmount := float64(0.00)

	for _, v := range cart.CartItems {
		totalAmount = totalAmount + v.GetAmount()
	}
	cart.TotalAmount = totalAmount
	return totalAmount
}
