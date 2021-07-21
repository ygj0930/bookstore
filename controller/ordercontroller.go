package controller

import (
	"bookstore/dao"
	"html/template"
	"net/http"
	"strconv"
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

//发货
func DoSendOrder(w http.ResponseWriter, r *http.Request) {
	//提取参数
	orderId := r.FormValue("orderId")

	//修改数据库中订单状态
	err := dao.TakeOrder(orderId, 1)
	if err != nil {
		panic(err)
	}

	//刷新订单管理页面
	DoGetOrders(w, r)
}

//订单详情页面
func DoGetOrderInfo(w http.ResponseWriter, r *http.Request) {
	//提取参数
	orderId := r.FormValue("orderId")

	//查询订单项
	items, err := dao.GetOrderItemsByOrderID(orderId)
	if err != nil {
		panic(err)
	}

	//渲染模板
	t := template.Must(template.ParseFiles("views/pages/order/order_info.html"))
	t.Execute(w, items)
}

//我的订单页面
func DoGetMyOrder(w http.ResponseWriter, r *http.Request) {
	session, _ := dao.GetSessionByCookie(r)
	userId := session.UserId

	//获取我的订单
	orders, err := dao.GetMyOrder(userId)
	if err != nil {
		panic(err)
	}

	session.Orders = orders
	//渲染模板
	t := template.Must(template.ParseFiles("views/pages/order/order.html"))
	t.Execute(w, session)
}

//确认收货/退款退货
func DoTakeOrder(w http.ResponseWriter, r *http.Request) {
	//提取参数
	orderId := r.FormValue("orderId")
	state := r.FormValue("state")
	iState, _ := strconv.Atoi(state)

	//修改数据库中订单状态
	err := dao.TakeOrder(orderId, iState)
	if err != nil {
		panic(err)
	}

	//刷新我的订单页面
	DoGetMyOrder(w, r)
}
