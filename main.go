package main

import (
	"bookstore/controller"
	"net/http"
)

func main() {
	//静态资源访问设置
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("views/static/"))))
	http.Handle("/pages/", http.StripPrefix("/pages/", http.FileServer(http.Dir("views/pages/"))))

	//页面请求
	//首页
	http.HandleFunc("/", controller.IndexHandler)
	http.HandleFunc("/main", controller.IndexHandler)

	//用户登录注册注销
	http.HandleFunc("/toLogin", controller.LoginPageHandler)
	http.HandleFunc("/toRegister", controller.RegisterPageHandler)
	http.HandleFunc("/logout", controller.LogoutHandler)

	//后台管理
	http.HandleFunc("/toManager", controller.ManagerPageHandler)

	//书籍管理
	http.HandleFunc("/getBooks", controller.BooksManagerPageHandler)
	http.HandleFunc("/getPageBooks", controller.PageBooksManagerPageHandler)
	http.HandleFunc("/getPageBooksByPrice", controller.PageBooksByPriceManagerPageHandler)
	http.HandleFunc("/toAddBook", controller.BookAddPageHandler)
	http.HandleFunc("/toUpdateBookPage", controller.BookUpdatePageHandler)

	//购物车页面
	http.HandleFunc("/getCartInfo", controller.CartPageHandler)

	//处理请求
	//登录注册相关请求
	http.HandleFunc("/login", controller.DoLogin)
	http.HandleFunc("/register", controller.DoRegister)
	http.HandleFunc("/findUserByName", controller.DoFindUserByName)

	//图书相关请求
	http.HandleFunc("/addBook", controller.DoAddBook)
	http.HandleFunc("/updateBook", controller.DoUpdateBook)
	http.HandleFunc("/deleteBook", controller.DoDeleteBook)

	//购物车相关请求
	http.HandleFunc("/AddBook2Cart", controller.DoAddBook2Cart)
	http.HandleFunc("/deleteCart", controller.DoDeleteCart)
	http.HandleFunc("/deleteCartItem", controller.DoDeleteCartItem)
	http.HandleFunc("/updateCartItemByGet", controller.DoUpdateCartItemByGET)
	http.HandleFunc("/updateCartItemByAjax", controller.DoUpdateCartItemByAJAX)

	//订单相关
	http.HandleFunc("/checkout", controller.DoCheckout)
	http.HandleFunc("/getOrders", controller.DoGetOrders)
	http.HandleFunc("/getOrderInfo", controller.DoGetOrderInfo)
	http.HandleFunc("/getMyOrder", controller.DoGetMyOrder)
	http.HandleFunc("/takeOrder", controller.DoTakeOrder)
	http.HandleFunc("/sendOrder", controller.DoSendOrder)

	//服务器启动
	http.ListenAndServe(":8080", nil)
}
