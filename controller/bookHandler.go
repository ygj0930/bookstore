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
	//获取数据
	Books,_ := dao.GetBooks()

	//渲染模板
	t := template.Must(template.ParseFiles("views/pages/manager/book_manager.html"))
	t.Execute(w,Books)
}

//跳转图书管理页面：分页查询
func PageBooksManagerPageHandler(w http.ResponseWriter,r *http.Request){
	//获取请求参数
	pageNoStr := r.FormValue("PageNo")
	pageNo,errParse := strconv.ParseInt(pageNoStr,10,64)
	if pageNoStr == "" || errParse != nil{
		pageNo = 1
	}
	//获取数据
	page,_ := dao.GetPageBooks(pageNo)

	//渲染模板
	t := template.Must(template.ParseFiles("views/pages/manager/book_manager.html"))

	//易错点：由于需要在template中调用方法，因此要传递指针类型
	t.Execute(w,&page)
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

//跳转修改图书页面
func BookUpdatePageHandler(w http.ResponseWriter,r *http.Request){
	//获取参数
	idStr := r.FormValue("bookId")
	bookId,_ := strconv.Atoi(idStr)

	//查询书籍信息
	book,err := dao.GetBookByID(bookId)
	if err!=nil {
		fmt.Println("GetBookByID error:",err)
	}

	//渲染模板
	t := template.Must(template.ParseFiles("views/pages/manager/book_edit.html"))
	t.Execute(w,book)
}
//修改图书
func DoUpdateBook(w http.ResponseWriter,r *http.Request){
	//获取参数
	idStr := r.PostFormValue("bookId")
	title := r.PostFormValue("title")
	price := r.PostFormValue("price")
	author := r.PostFormValue("author")
	sales := r.PostFormValue("sales")
	stock := r.PostFormValue("stock")
	img_path := r.PostFormValue("img_path")

	//类型转换
	iId,_ := strconv.Atoi(idStr)
	fPrice,_ := strconv.ParseFloat(price,2)
	iSales,_ := strconv.Atoi(sales)
	iStock,_ := strconv.Atoi(stock)
	//进行插入
	errIns := dao.UpdateBook(&model.Book{
		ID:   iId,
		Title:   title,
		Author:  author,
		Price:   fPrice,
		Sales:   iSales,
		Stock:   iStock,
		ImgPath: img_path,
	})
	if errIns != nil {
		fmt.Println("DoUpdateBook error:",errIns)
	}

	//内部跳转到图书管理页面
	BooksManagerPageHandler(w,r)
}