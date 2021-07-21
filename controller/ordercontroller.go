package controller

import (
	"bookstore/dao"
	"html/template"
	"net/http"
)

//订单管理页面
func DoGetOrders(w http.ResponseWriter, r *http.Request) {
	//获取所有订单
	orders, err := dao.GetOrders()
	if err != nil {
		panic(err)
	}

	//渲染模板
	t := template.Must(template.ParseFiles("views/pages/order/order_manager.html"))
	t.Execute(w, orders)
}
