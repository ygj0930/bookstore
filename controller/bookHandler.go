package controller

import (
	"bookstore/dao"
	"bookstore/model"
	"fmt"
	"html/template"
	"net/http"
	"strconv"
)

//跳转图书管理页面
func BooksManagerPageHandler(w http.ResponseWriter,r *http.Request){
	//渲染模板
	t := template.Must(template.ParseFiles("views/pages/manager/book_manager.html"))

	//获取数据
	Books,_ := dao.GetBooks()
	t.Execute(w,Books)
}

//跳转添加图书页面
func BookAddPageHandler(w http.ResponseWriter,r *http.Request){
	//渲染模板
	t := template.Must(template.ParseFiles("views/pages/manager/book_add.html"))
	t.Execute(w,"")
}
//新增图书
func DoAddBook(w http.ResponseWriter,r *http.Request){
	//获取参数
	title := r.PostFormValue("title")
	price := r.PostFormValue("price")
	author := r.PostFormValue("author")
	sales := r.PostFormValue("sales")
	stock := r.PostFormValue("stock")
	img_path := r.PostFormValue("img_path")

	//类型转换
	fPrice,_ := strconv.ParseFloat(price,2)
	iSales,_ := strconv.Atoi(sales)
	iStock,_ := strconv.Atoi(stock)
	//进行插入
	errIns := dao.AddBook(&model.Book{
		Title:   title,
		Author:  author,
		Price:   fPrice,
		Sales:   iSales,
		Stock:   iStock,
		ImgPath: img_path,
	})
	if errIns != nil {
		fmt.Println("DoAddBook error:",errIns)
	}

	//内部跳转到图书管理页面
	BooksManagerPageHandler(w,r)
}

//删除图书
func DoDeleteBook(w http.ResponseWriter,r *http.Request){
	//获取参数
	idStr := r.FormValue("bookId")
	bookId,_ := strconv.Atoi(idStr)

	//进行删除
	err := dao.DeleteBook(bookId)
	if err!=nil {
		fmt.Println("DoDeleteBook error:",err)
	}

	//内部跳转回图书管理页面
	BooksManagerPageHandler(w,r)
}