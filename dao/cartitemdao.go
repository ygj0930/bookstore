package dao

import (
	"bookstore/model"
	"bookstore/utils"
	"fmt"
	"strconv"
)

//新增购物项
func AddCartItem(cartItem *model.CartItem) error {
	sqlStr := "insert into cart_items(count,amount,book_id,cart_id) values(?,?,?,?)"

	_, errIns := utils.Db.Exec(sqlStr, cartItem.Count, cartItem.GetAmount(), cartItem.Book.ID, cartItem.CartId)
	if errIns != nil {
		return errIns
	}
	return nil
}

//删除购物项
func DeleteCartItem(cartItemId string) error {
	sqlStr := "delete from cart_items where id = ?"

	_, errDel := utils.Db.Exec(sqlStr, cartItemId)
	if errDel != nil {
		return errDel
	}
	return nil
}

//修改购物项
func UpdateCartItem(cartItem *model.CartItem) error {
	sqlStr := "update cart_items set count=?,amount=? where id=?"
	_, errUpd := utils.Db.Exec(sqlStr, cartItem.Count, cartItem.GetAmount(), cartItem.ID)
	if errUpd != nil {
		return errUpd
	}
	return nil
}

//根据购物车id查询所有购物项
func GetCartItemsByCartID(cart_id string) ([]*model.CartItem, error) {
	sqlStr := "select id,count,amount,book_id from cart_items where cart_id=?"
	rows, errQuery := utils.Db.Query(sqlStr, cart_id)
	if errQuery != nil {
		return nil, errQuery
	}
	cartItems := make([]*model.CartItem, 0)
	for rows.Next() {
		cartItem := &model.CartItem{}
		var ibookId int
		errScan := rows.Scan(&cartItem.ID, &cartItem.Count, &cartItem.Amount, &ibookId)
		if errScan != nil {
			fmt.Println("GetCartItemsByCartID Scan err:", errScan)
			return nil, errScan
		}
		book, errBook := GetBookByID(ibookId)
		if errBook != nil || book == nil {
			fmt.Println("GetCartItemsByCartID Scan err:", errBook)
			return nil, errBook
		}
		cartItem.Book = book
		cartItem.CartId = cart_id

		cartItems = append(cartItems, cartItem)
	}
	return cartItems, nil
}

//根据购物车id和书籍id查询具体购物项
func GetCartItemByCartIDAndBookID(cart_id, book_id string) (*model.CartItem, error) {
	sqlStr := "select id,count,amount from cart_items where book_id=? and cart_id=?"
	row := utils.Db.QueryRow(sqlStr, book_id, cart_id)
	cartItem := &model.CartItem{}
	errScan := row.Scan(&cartItem.ID, &cartItem.Count, &cartItem.Amount)
	if errScan != nil {
		fmt.Println("GetCartItemByCartIDAndBookID Scan err:", errScan)
		return nil, errScan
	}
	ibookId, _ := strconv.Atoi(book_id)
	book, errBook := GetBookByID(ibookId)
	if errBook != nil || book == nil {
		fmt.Println("GetCartItemByCartIDAndBookID Scan err:", errBook)
		return nil, errBook
	}
	cartItem.Book = book
	cartItem.CartId = cart_id
	return cartItem, nil
}
