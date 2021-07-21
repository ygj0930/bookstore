package dao

import (
	"bookstore/model"
	"bookstore/utils"
	"fmt"
)

//新增订单项
func AddOrderItem(orderItem *model.OrderItem) error {
	sqlStr := "insert into order_items(count,amount,title,author,price,img_path,order_id) values(?,?,?,?,?,?,?)"

	_, errIns := utils.Db.Exec(sqlStr, orderItem.Count, orderItem.Amount, orderItem.Title, orderItem.Author, orderItem.Price, orderItem.ImgPath, orderItem.OrderID)
	if errIns != nil {
		return errIns
	}
	return nil
}

//查询某订单的所有订单项
func GetOrderItemsByOrderID(orderId string) ([]*model.OrderItem, error) {
	sqlStr := "select * from order_items where order_id=?"
	rows, errQuery := utils.Db.Query(sqlStr, orderId)
	if errQuery != nil {
		return nil, errQuery
	}

	orderItems := make([]*model.OrderItem, 0)
	for rows.Next() {
		orderItem := &model.OrderItem{}
		errScan := rows.Scan(&orderItem.OrderItemID, &orderItem.Count, &orderItem.Amount, &orderItem.Title, &orderItem.Author, &orderItem.Price, &orderItem.ImgPath, &orderItem.OrderID)
		if errScan != nil {
			fmt.Println("GetOrderItemsByOrderID Scan err:", errScan)
			return nil, errScan
		}
		orderItems = append(orderItems, orderItem)
	}
	return orderItems, nil
}
