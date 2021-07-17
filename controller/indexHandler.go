package controller

import (
	"html/template"
	"net/http"
)

//首页请求处理：渲染并返回首页
func IndexHandler(w http.ResponseWriter,r *http.Request){
	//渲染模板
	t := template.Must(template.ParseFiles("views/index.html"))

	t.Execute(w,"")
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


