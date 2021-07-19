package dao

import (
	"bookstore/model"
	"bookstore/utils"
	"fmt"
)

//新增购物车
func AddCart(cart *model.Cart) error {
	//保存表头
	sqlStr := "insert into carts values(?,?,?,?,?)"
	_, errIns := utils.Db.Exec(sqlStr, cart.ID, cart.GetTotalCount(), cart.GetTotalAmount(), cart.UserID, cart.SessionID)
	if errIns != nil {
		return errIns
	}

	//保存购物项明细
	cartItems := cart.CartItems
	for _, v := range cartItems {
		AddCartItem(v)
	}
	return nil
}

//删除购物车
func DeleteCart(cartId string) error {
	//删除明细
	itemErr := DeleteCartItemsByCartId(cartId)
	if itemErr != nil {
		return itemErr
	}

	//删除购物车
	sqlStr := "delete from carts where id = ?"
	_, errDel := utils.Db.Exec(sqlStr, cartId)
	if errDel != nil {
		return errDel
	}
	return nil
}

//根据sessionId查询用户当前会话期的购物车
func GetCartBySessionId(sessionId string) (*model.Cart, error) {
	sqlStr := "select * from carts where session_id = ?"
	row := utils.Db.QueryRow(sqlStr, sessionId)
	cart := &model.Cart{}
	errScan := row.Scan(&cart.ID, &cart.TotalCount, &cart.TotalAmount, &cart.UserID, &cart.SessionID)
	if errScan != nil {
		fmt.Println("GetCartBySessionId Scan err:", errScan)
		return nil, errScan
	}
	cartItems, errItems := GetCartItemsByCartID(cart.ID)
	if errItems != nil {
		fmt.Println("GetCartBySessionId Scan err:", errItems)
		return nil, errItems
	}
	cart.CartItems = cartItems
	return cart, nil
}

func UpdateCart(cart *model.Cart) error {
	sqlStr := "update carts set total_count=?,total_amount=? where id=?"
	_, errUpd := utils.Db.Exec(sqlStr, cart.GetTotalCount(), cart.GetTotalAmount(), cart.ID)
	if errUpd != nil {
		return errUpd
	}
	return nil
}
