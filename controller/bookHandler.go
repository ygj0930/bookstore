package controller

import (
	"bookstore/dao"
	"html/template"
	"net/http"
)

//跳转图书管理页面
func BooksManagerPageHandler(w http.ResponseWriter,r *http.Request){
	//渲染模板
	t := template.Must(template.ParseFiles("views/pages/manager/book_manager.html"))

	//获取数据
	Books,_ := dao.GetBooks()
	t.Execute(w,Books)
}
