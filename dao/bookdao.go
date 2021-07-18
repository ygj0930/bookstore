package dao

import (
	"bookstore/model"
	"bookstore/utils"
	"fmt"
)

//获取所有图书
func GetBooks()([]*model.Book,error) {
	sqlStr := "select id,title,author,price,sales,stock,img_path from books"

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
//分页查询图书
func GetPageBooks(pageNo int64) (*model.Page,error) {
	//查询总记录数
	countSql := "select count(*) from books"
	var recordTotal int64
	countRow := utils.Db.QueryRow(countSql)
	countRow.Scan(&recordTotal)

	//设置每页大小
	var pageSize int64 = 10

	//计算总页数
	var pageTotal int64
	if recordTotal % pageSize == 0 {
		pageTotal = recordTotal / pageSize
	}else{
		pageTotal = recordTotal / pageSize + 1
	}

	//查询当前页记录
	booksSql := "select * from books limit ?,?"
	rows,_ := utils.Db.Query(booksSql,(pageNo-1)*pageSize,pageSize)
	var books []*model.Book
	for rows.Next() {
		book := &model.Book{}
		errScan := rows.Scan(&book.ID,&book.Title,&book.Author,&book.Price,&book.Sales,&book.Stock,&book.ImgPath)
		if errScan != nil {
			continue
		}
		books = append(books,book)
	}

	//创建分页记录返回
	pageInfo := model.Page{
		Books:       books,
		PageNo:      pageNo,
		PageSize:    pageSize,
		PageTotal:   pageTotal,
		RecordTotal: recordTotal,
	}
	return &pageInfo,nil
}


//分页查询价格范围图书
func GetPageBooksByPrice(pageNo int64,minPrice string,maxPrice string) (*model.Page,error) {
	//查询总记录数
	countSql := "select count(*) from books where price between ? and ?"
	var recordTotal int64
	countRow := utils.Db.QueryRow(countSql,minPrice,maxPrice)
	countRow.Scan(&recordTotal)

	//设置每页大小
	var pageSize int64 = 10

	//计算总页数
	var pageTotal int64
	if recordTotal % pageSize == 0 {
		pageTotal = recordTotal / pageSize
	}else{
		pageTotal = recordTotal / pageSize + 1
	}

	//查询当前页记录
	booksSql := "select * from books where price between ? and ? limit ?,?"
	rows,_ := utils.Db.Query(booksSql,minPrice,maxPrice,(pageNo-1)*pageSize,pageSize)
	var books []*model.Book
	for rows.Next() {
		book := &model.Book{}
		errScan := rows.Scan(&book.ID,&book.Title,&book.Author,&book.Price,&book.Sales,&book.Stock,&book.ImgPath)
		if errScan != nil {
			continue
		}
		books = append(books,book)
	}

	//创建分页记录返回
	pageInfo := model.Page{
		Books:       books,
		PageNo:      pageNo,
		PageSize:    pageSize,
		PageTotal:   pageTotal,
		RecordTotal: recordTotal,
		MinPrice: minPrice,
		MaxPrice: maxPrice,
	}
	return &pageInfo,nil
}

//查询某本图书
func GetBookByID(bookID int) (*model.Book,error){
	sqlStr := "select id,title,author,price,sales,stock,img_path from books where id = ?"

	row := utils.Db.QueryRow(sqlStr,bookID)
	book := &model.Book{}
	errScan :=row.Scan(&book.ID,&book.Title,&book.Author,&book.Price,&book.Sales,&book.Stock,&book.ImgPath)
	if errScan!=nil {
		fmt.Println("GetBookByID Scan err:",errScan)
		return nil,errScan
	}
	return book,nil
}

//新增书籍
func AddBook(b *model.Book) error {
	sqlStr := "insert into books(title,author,price,sales,stock,img_path) values(?,?,?,?,?,?)"

	_,errIns := utils.Db.Exec(sqlStr,b.Title,b.Author,b.Price,b.Sales,b.Stock,b.ImgPath)
	if errIns != nil {
		return errIns
	}
	return nil
}

//修改书籍
func UpdateBook(b *model.Book) error {
	sqlStr := "update books set title = ?,author=?,price=?,sales=?,stock=?,img_path=? where id = ?"

	_,errUpd := utils.Db.Exec(sqlStr,b.Title,b.Author,b.Price,b.Sales,b.Stock,b.ImgPath,b.ID)
	if errUpd != nil {
		return errUpd
	}
	return nil
}
//删除书籍
func DeleteBook(bookId int) error {
	sqlStr := "delete from books where id = ?"

	_,errDel := utils.Db.Exec(sqlStr,bookId)

	if errDel != nil {
		return errDel
	}
	return nil
}