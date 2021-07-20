package controller

import (
	"bookstore/dao"
	"bookstore/model"
	"bookstore/utils"
	"html/template"
	"net/http"
	"strconv"
	"time"
)

//查看购物车
func CartPageHandler(w http.ResponseWriter, r *http.Request) {
	//查询当前用户的购物车
	session, _ := dao.GetSessionByCookie(r)
	cart, _ := dao.GetCartBySessionId(session.SessionId)
	session.Cart = cart
	t := template.Must(template.ParseFiles("views/pages/cart/cart.html"))
	t.Execute(w, session)
}

//清空购物车
func DoDeleteCart(w http.ResponseWriter, r *http.Request) {
	cartId := r.FormValue("cartId")

	//删除数据库购物车以及明细
	dao.DeleteCart(cartId)

	//删除会话中的购物车
	session, _ := dao.GetSessionByCookie(r)
	session.Cart = nil

	//跳转回购物车页面
	CartPageHandler(w, r)
}

//删除购物项
func DoDeleteCartItem(w http.ResponseWriter, r *http.Request) {
	cartItemId := r.FormValue("cartItemId")
	iCartItemID, _ := strconv.ParseInt(cartItemId, 10, 64)
	//删除数据库中购物项
	dao.DeleteCartItem(cartItemId)

	//更新购物车
	session, _ := dao.GetSessionByCookie(r)
	cart, _ := dao.GetCartBySessionId(session.SessionId)
	if cart != nil {
		cartItems := cart.CartItems
		for index, item := range cartItems {
			//寻找要删除的购物项
			if item.ID == iCartItemID {
				//进行删除：从切片里剔除
				newCartItems := append(cartItems[:index], cartItems[index+1:]...)
				//更新购物车购物项
				cart.CartItems = newCartItems
			}
		}
	}
	dao.UpdateCart(cart)

	//跳转回购物车页面
	CartPageHandler(w, r)
}

//修改购物项
func DoUpdateCartItem(w http.ResponseWriter, r *http.Request) {
	cartId := r.FormValue("cartId")
	bookId := r.FormValue("bookId")
	bookCount := r.FormValue("bookCount")
	iBookCount, _ := strconv.ParseInt(bookCount, 10, 64)

	//从数据库查询购物项
	cartItem, _ := dao.GetCartItemByCartIDAndBookID(cartId, bookId)
	//修改购物项数量
	cartItem.Count = iBookCount
	cartItem.Amount = cartItem.GetAmount()

	//更新数据库购物项
	dao.UpdateCartItem(cartItem)

	//更新购物车
	session, _ := dao.GetSessionByCookie(r)
	cart, _ := dao.GetCartBySessionId(session.SessionId)
	if cart != nil {
		cartItems := cart.CartItems
		for _, item := range cartItems {
			//寻找要更新的购物项
			if item.ID == cartItem.ID {
				item.Count = iBookCount
				item.Amount = item.GetAmount()
			}
		}
	}
	dao.UpdateCart(cart)

	//跳转回购物车页面
	CartPageHandler(w, r)
}

//加入购物车
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
			if cartItem.Count > int64(cartItem.Book.Stock) {
				res = "库存不足!"
			} else {
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
				if cartItem.Count > int64(cartItem.Book.Stock) {
					res = "库存不足!"
				} else {
					//往数据库插入购物项
					dao.AddCartItem(cartItem)

					//往内存中购物车增加购物项
					cart.CartItems = append(cart.CartItems, cartItem)
				}
			} else { //有购物项，则增加数量
				newCount := cartItem.Count + 1
				if newCount > int64(cartItem.Book.Stock) {
					res = "库存不足!"
				} else {
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
			}
			//更新购物车的总金额总数量
			dao.UpdateCart(cart)
		}
		if res != "库存不足!" {
			countStr := strconv.FormatInt(cart.GetTotalCount(), 10)
			res = "您刚刚购买了《" + book.Title + "》,购物车中已有 " + countStr + " 本图书"
		}
	}

	w.Write([]byte(res))
}

//结账：将当前会话中的购物车转化为订单，并清空购物车信息[数据库层面、内存层面]
func DoCheckout(w http.ResponseWriter, r *http.Request) {
	//获取本次会话中的购物车
	session, _ := dao.GetSessionByCookie(r)
	cart, _ := dao.GetCartBySessionId(session.SessionId)

	//===拼接订单===
	//创建订单ID
	orderId := utils.CreateUUID()
	//创建订单项
	var orderItems []*model.OrderItem
	for _, ct := range cart.CartItems {
		book := ct.Book

		orderItem := &model.OrderItem{
			Count:   ct.Count,
			Amount:  ct.Amount,
			Title:   book.Title,
			Author:  book.Author,
			Price:   book.Price,
			ImgPath: book.ImgPath,
			OrderID: orderId,
		}
		orderItems = append(orderItems, orderItem)
		//更新图书库存和销量
		book.Sales = int(int64(book.Sales) + ct.Count)
		book.Stock = int(int64(book.Stock) - ct.Count)
		dao.UpdateBook(book)
	}
	//创建订单
	order := &model.Order{
		OrderID:     orderId,
		CreateTime:  time.Now().Format("2006-01-02 15:04:05"),
		TotalCount:  cart.TotalCount,
		TotalAmount: cart.TotalAmount,
		State:       0,
		UserID:      session.UserId,
		OrderItems:  orderItems,
	}
	//插入数据库
	dao.AddOrder(order)

	//删除购物车
	dao.DeleteCart(cart.ID)
	session.Cart = nil

	//跳转结算页面
	//渲染模板
	t := template.Must(template.ParseFiles("views/pages/cart/checkout.html"))
	session.OrderId = orderId
	t.Execute(w, session)
}
