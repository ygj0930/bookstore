package dao

import (
	"bookstore/utils"
	"bookstore/model"
)

//获取所有图书
func GetBooks()([]*model.Book,error) {
	sqlStr := "select id , title,author ,price ,sales ,stock ,img_path  from books"

	rows,err := utils.Db.Query(sqlStr)
	if err != nil {
		return nil, err
	}
	var books []*model.Book
	for rows.Next() {
		book := &model.Book{}
		errScan := rows.Scan(&book.ID,&book.Title,&book.Author,&book.Price,&book.Sales,&book.Stock,&book.ImgPath)
		if errScan != nil {
			continue
		}
		books = append(books,book)
	}
	return books,nil
}
