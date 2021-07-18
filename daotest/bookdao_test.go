package daotest

import (
	"bookstore/dao"
	"bookstore/model"
	"testing"
)

func TestGetBooks(t *testing.T) {
	res,err := dao.GetBooks()
	if err != nil {
		t.Error(err)
	}
	for _,book := range res {
		t.Log("TestGetBook-book:",book)
	}
}

func TestGetPageBooks(t *testing.T) {
	res,err := dao.GetPageBooks(1)
	if err != nil {
		t.Error(err)
	}
	t.Log("TestGetPageBooks-pageNo:",res.PageNo)
	t.Log("TestGetPageBooks-pageTotal:",res.PageTotal)
	t.Log("TestGetPageBooks-recTotal:",res.RecordTotal)
	for _,v := range res.Books {
		t.Log("TestGetPageBooks-book:",v)
	}
}

func TestGetPageBooksByPrice(t *testing.T) {
	res,err := dao.GetPageBooksByPrice(1,"10","40")
	if err != nil {
		t.Error(err)
	}
	t.Logf("GetPageBooksByPrice:%+v",res)
}

func TestAddBook(t *testing.T) {
	book := model.Book{
		Title: "测试添加",
		Author: "测试添加",
		Price: 10,
		Sales: 1,
		Stock: 99,
		ImgPath: "static/img/default.jpg",
	}
	err := dao.AddBook(&book)
	if err!=nil {
		t.Error(err)
	}
}

//修改图书测试
func TestUpdateBook(t *testing.T) {
	book := model.Book{
		ID: 7,
		Title: "测试修改",
		Author: "测试修改",
		Price: 10,
		Sales: 1,
		Stock: 99,
		ImgPath: "nil",
	}
	err := dao.UpdateBook(&book)
	if err!=nil {
		t.Error(err)
	}
}

func TestDeleteBook(t *testing.T){
	bookId := 45
	err := dao.DeleteBook(bookId)
	if err!=nil {
		t.Error(err)
	}
}

func TestGetBookByID(t *testing.T){
	bookId := 1
	book,err := dao.GetBookByID(bookId)
	if err!=nil {
		t.Error(err)
	}else{
		t.Log(book)
	}
}