package daotest

import (
	"bookstore/dao"
	"bookstore/model"
	"bookstore/utils"
	"testing"
)

func TestDeleteCart(t *testing.T) {
	err := dao.DeleteCart("c99bd544-6ac5-4522-5079-9ea860b170a6")
	if err != nil {
		t.Log(err)
	}
}

func TestGetCartItemsByCartID(t *testing.T) {
	res, _ := dao.GetCartItemsByCartID("0ab18110-d15f-4458-5027-bb03840f16cd")
	for _, v := range res {
		t.Logf("items:%+v", v)
	}
}

func TestGetCartItemByCartIDAndBookID(t *testing.T) {
	res, _ := dao.GetCartItemByCartIDAndBookID("0ab18110-d15f-4458-5027-bb03840f16cd", "3")
	t.Logf("items:%+v", res)
}

func TestGetCartByUserId(t *testing.T) {
	res, _ := dao.GetCartBySessionId("1")
	t.Logf("cart:%+v", res)
}

func TestAddCart(t *testing.T) {
	book1 := &model.Book{
		ID:    1,
		Price: 27,
	}
	book2 := &model.Book{
		ID:    2,
		Price: 232,
	}
	book3 := &model.Book{
		ID:    3,
		Price: 44,
	}

	cartId := utils.CreateUUID()

	cartItem1 := &model.CartItem{
		Book:   book1,
		Count:  1,
		CartId: cartId,
	}
	cartItem2 := &model.CartItem{
		Book:   book2,
		Count:  2,
		CartId: cartId,
	}
	cartItem3 := &model.CartItem{
		Book:   book3,
		Count:  3,
		CartId: cartId,
	}

	cartItems := make([]*model.CartItem, 0)
	cartItems = append(cartItems, cartItem1, cartItem2, cartItem3)

	cart := model.Cart{
		ID:        cartId,
		CartItems: cartItems,
		UserID:    1,
		SessionID: "1",
	}
	dao.AddCart(&cart)
}
