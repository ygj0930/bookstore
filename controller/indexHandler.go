package controller

import (
	"bookstore/dao"
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

	//渲染模板
	t := template.Must(template.ParseFiles("views/index.html"))

	t.Execute(w,page)
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


