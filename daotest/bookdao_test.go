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

func TestDeleteBook(t *testing.T){
	bookId := 45
	err := dao.DeleteBook(bookId)
	if err!=nil {
		t.Error(err)
	}
}