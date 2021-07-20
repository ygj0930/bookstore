package daotest

import (
	"bookstore/dao"
	"bookstore/model"
	"bookstore/utils"
	"testing"
	"time"
)

func TestAddOrder(t *testing.T) {
	orderId := utils.CreateUUID()

	orderItem1 := &model.OrderItem{
		Count:   1,
		Amount:  10,
		Title:   "书本1",
		Author:  "作者1",
		Price:   10,
		ImgPath: "nil",
		OrderID: orderId,
	}
	orderItem2 := &model.OrderItem{
		Count:   2,
		Amount:  20,
		Title:   "书本2",
		Author:  "作者2",
		Price:   10,
		ImgPath: "nil",
		OrderID: orderId,
	}
	orderItems := make([]*model.OrderItem, 0)
	orderItems = append(orderItems, orderItem1, orderItem2)
	order := &model.Order{
		OrderID:     orderId,
		CreateTime:  time.Now().Format("2006-01-02 15:04:05"),
		TotalCount:  3,
		TotalAmount: 30,
		State:       0,
		UserID:      1,
		OrderItems:  orderItems,
	}
	err := dao.AddOrder(order)
	if err != nil {
		t.Log(err)
	}
}
