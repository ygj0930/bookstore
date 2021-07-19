package controller

import (
	"bookstore/dao"
	"fmt"
	"html/template"
	"net/http"
	"strconv"
)

//首页请求处理：渲染并返回首页
func IndexHandler(w http.ResponseWriter,r *http.Request){
	//获取请求参数
	pageNoStr := r.FormValue("PageNo")
	pageNo,errParse := strconv.ParseInt(pageNoStr,10,64)
	if pageNoStr == "" || errParse != nil{
		pageNo = 1
	}
	//获取数据
	page,_ := dao.GetPageBooks(pageNo)

	//获取会话信息
	session,err := dao.GetSessionByCookie(r)
	if err != nil {
		fmt.Println("session not exit！")
	}else{
		page.IsLogin = true
		page.UserName = session.UserName
	}

	//渲染模板
	t := template.Must(template.ParseFiles("views/index.html"))

	t.Execute(w,page)
}

//首页价格查询请求处理：渲染并返回首页
func PageBooksByPriceManagerPageHandler(w http.ResponseWriter,r *http.Request){
	//获取请求参数
	pageNoStr := r.FormValue("PageNo")
	minPriceStr := r.FormValue("minPrice")
	maxPriceStr := r.FormValue("maxPrice")
	pageNo,errParse := strconv.ParseInt(pageNoStr,10,64)
	if pageNoStr == "" || errParse != nil{
		pageNo = 1
	}
	fmt.Println("pageNoStr:",pageNoStr,"minPriceStr:",minPriceStr,"maxPriceStr:",maxPriceStr)
	if minPriceStr=="" && maxPriceStr==""{
		//获取数据
		page,_ := dao.GetPageBooks(pageNo)
		//获取会话信息
		session,err := dao.GetSessionByCookie(r)
		if err != nil {
			fmt.Println("session not exit！")
		}else{
			page.IsLogin = true
			page.UserName = session.UserName
		}

		//渲染模板
		t := template.Must(template.ParseFiles("views/index.html"))

		t.Execute(w,page)
	}else{
		//获取数据
		page,_ := dao.GetPageBooksByPrice(pageNo,minPriceStr,maxPriceStr)
		//获取会话信息
		session,err := dao.GetSessionByCookie(r)
		if err != nil {
			fmt.Println("session not exit！")
		}else{
			page.IsLogin = true
			page.UserName = session.UserName
		}
		//渲染模板
		t := template.Must(template.ParseFiles("views/searchOfPrice.html"))

		t.Execute(w,page)
	}
}


//跳转登录页面
func LoginPageHandler(w http.ResponseWriter,r *http.Request){
	//渲染模板
	t := template.Must(template.ParseFiles("views/pages/user/login.html"))

	t.Execute(w,"")
}

//跳转注册页面
func RegisterPageHandler(w http.ResponseWriter,r *http.Request){
	//渲染模板
	t := template.Must(template.ParseFiles("views/pages/user/regist.html"))

	t.Execute(w,"")
}


