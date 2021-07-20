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
