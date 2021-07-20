package dao

import (
	"bookstore/model"
	"bookstore/utils"
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
