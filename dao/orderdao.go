package dao

import (
	"bookstore/model"
	"bookstore/utils"
)

//新增订单
func AddOrder(order *model.Order) error {
	//保存表头
	sqlStr := "insert into orders values(?,?,?,?,?,?)"
	_, errIns := utils.Db.Exec(sqlStr, order.OrderID, order.CreateTime, order.TotalCount, order.TotalAmount, order.State, order.UserID)
	if errIns != nil {
		return errIns
	}

	//保存购物项明细
	orderItems := order.OrderItems
	for _, v := range orderItems {
		AddOrderItem(v)
	}
	return nil
}

//订单管理：获取所有订单
func GetOrders() ([]*model.Order, error) {
	sqlStr := "select * from orders"
	orderRows, errRows := utils.Db.Query(sqlStr)
	if errRows != nil {
		return nil, errRows
	}

	var orders []*model.Order
	for orderRows.Next() {
		order := &model.Order{}
		errScan := orderRows.Scan(&order.OrderID, &order.CreateTime, &order.TotalCount, &order.TotalAmount, &order.State, &order.UserID)
		if errScan != nil {
			return nil, errScan
		}
		orderItems, errItems := GetOrderItemsByOrderID(order.OrderID)
		if errItems != nil {
			return nil, errItems
		}
		order.OrderItems = orderItems

		orders = append(orders, order)
	}

	return orders, nil
}
