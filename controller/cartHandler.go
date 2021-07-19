package controller

import (
	"bookstore/dao"
	"bookstore/model"
	"bookstore/utils"
	"html/template"
	"net/http"
	"strconv"
)

func CartPageHandler(w http.ResponseWriter, r *http.Request) {
	//查询当前用户的购物车
	session, _ := dao.GetSessionByCookie(r)
	cart, _ := dao.GetCartBySessionId(session.SessionId)
	session.Cart = cart
	t := template.Must(template.ParseFiles("views/pages/cart/cart.html"))
	t.Execute(w, session)
}

func DoAddBook2Cart(w http.ResponseWriter, r *http.Request) {
	res := "请先进行登录再操作！"
	//判断是否登录
	if flag := dao.IsLogin(r); flag {

		//提取参数
		bookId := r.PostFormValue("bookId")
		//获取要购买的图书
		iBookId, _ := strconv.Atoi(bookId)
		book, _ := dao.GetBookByID(iBookId)

		//查询当前用户的购物车
		//从会话中提取user_id
		session, _ := dao.GetSessionByCookie(r)
		userId := session.UserId
		cart, errCart := dao.GetCartBySessionId(session.SessionId)

		//无车则新增
		if cart == nil || errCart != nil {
			cartId := utils.CreateUUID()

			cartItems := make([]*model.CartItem, 0)
			cartItem := &model.CartItem{
				Book:   book,
				Count:  1,
				Amount: book.Price,
				CartId: cartId,
			}
			cartItems = append(cartItems, cartItem)

			cart = &model.Cart{
				ID:        cartId,
				CartItems: cartItems,
				UserID:    userId,
				SessionID: session.SessionId,
			}
			cartAddErr := dao.AddCart(cart)
			if cartAddErr != nil {
				res = "创建购物车失败！"
			}
		} else { //有车

			//检查购物车中是否已有重复图书购买项
			cartItem, errItem := dao.GetCartItemByCartIDAndBookID(cart.ID, bookId)
			if cartItem == nil || errItem != nil { //无购物车项，则新增购物项
				cartItem = &model.CartItem{
					Book:   book,
					Count:  1,
					CartId: cart.ID,
				}
				//往数据库插入购物项
				dao.AddCartItem(cartItem)

				//往内存中购物车增加购物项
				cart.CartItems = append(cart.CartItems, cartItem)

			} else { //有购物项，则增加数量
				newCount := cartItem.Count + 1
				newAmount := cartItem.GetAmount()

				//修改数据库中对应购物项的数量和金额
				cartItem.Count = newCount
				cartItem.Amount = newAmount
				dao.UpdateCartItem(cartItem)

				//修改内存中购物车对应购物项的数量和金额
				for _, v := range cart.CartItems {
					if v.ID == cartItem.ID {
						v.Count = newCount
						v.Amount = newAmount
					}
				}
			}
			//更新购物车的总金额总数量
			dao.UpdateCart(cart)
		}
		countStr := strconv.FormatInt(cart.GetTotalCount(), 10)
		res = "您刚刚购买了《" + book.Title + "》,购物车中已有 " + countStr + " 本图书"
	}

	w.Write([]byte(res))
}
