package model

type CartItem struct {
	ID     int64
	Book   *Book
	Count  int64
	Amount float64
	CartId string //购物车id
}

func (cartItem *CartItem) GetAmount() float64 {
	return float64(cartItem.Count) * cartItem.Book.Price
}
