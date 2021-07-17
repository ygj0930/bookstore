package controller

import (
	"html/template"
	"net/http"
)

//跳转后台管理页面
func ManagerPageHandler(w http.ResponseWriter,r *http.Request){
	//渲染模板
	t := template.Must(template.ParseFiles("views/pages/manager/manager.html"))

	t.Execute(w,"")
}
