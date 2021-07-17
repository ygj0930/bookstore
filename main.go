package main

import (
	"bookstore/controller"
	"net/http"
)

func main() {
	//静态资源访问设置
	http.Handle("/static/",http.StripPrefix("/static/",http.FileServer(http.Dir("views/static/"))))
	http.Handle("/pages/",http.StripPrefix("/pages/",http.FileServer(http.Dir("views/pages/"))))

	//页面请求
	//首页
	http.HandleFunc("/main",controller.IndexHandler)

	//用户登录注册
	http.HandleFunc("/toLogin",controller.LoginPageHandler)
	http.HandleFunc("/toRegister",controller.RegisterPageHandler)

	//后台管理
	http.HandleFunc("/toManager",controller.ManagerPageHandler)
	http.HandleFunc("/toAddBook",controller.BookAddPageHandler)
	http.HandleFunc("/getBooks",controller.BooksManagerPageHandler)


	//处理请求
	//登录注册相关请求
	http.HandleFunc("/login",controller.DoLogin)
	http.HandleFunc("/register",controller.DoRegister)
	http.HandleFunc("/findUserByName",controller.DoFindUserByName)

	//图书相关请求
	http.HandleFunc("/addBook",controller.DoAddBook)


	//服务器启动
	http.ListenAndServe(":8080",nil)
}
