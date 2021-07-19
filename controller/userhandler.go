package controller

import (
	"bookstore/dao"
	"bookstore/model"
	"bookstore/utils"
	"fmt"
	"html/template"
	"net/http"
)

//处理登录请求
func DoLogin(w http.ResponseWriter,r *http.Request){
	//获取请求参数
	username := r.PostFormValue("username")
	password := r.PostFormValue("password")

	//进行验证
	user,err := dao.CheckUserNameAndPassword(username,password)
	if err != nil {
		fmt.Println("DoLogin error:",err)
	}
	if user != nil {
		//创建session
		sessionId := utils.CreateUUID()
		session := &model.Session{
			SessionId: sessionId,
			UserName:  username,
			UserId:    user.ID,
		}
		//保存
		_ = dao.AddSession(session)

		//设置cookie，回传给浏览器
		cookie := &http.Cookie{
			Name:     "sessionId",
			Value:    sessionId,
			HttpOnly: true,
		}
		http.SetCookie(w,cookie)

		t := template.Must(template.ParseFiles("views/pages/user/login_success.html"))
		t.Execute(w,session)
	}else{
		t := template.Must(template.ParseFiles("views/pages/user/login.html"))
		t.Execute(w,"登录失败，请检查输入的用户名和密码。")
	}
}

//处理注册请求
func DoRegister(w http.ResponseWriter,r *http.Request){
	//获取请求参数
	username := r.PostFormValue("username")
	password := r.PostFormValue("password")
	email := r.PostFormValue("email")

	//验证是否重复
	user,err := dao.CheckUserName(username)
	if err != nil {
		fmt.Println("DoRegister error:",err)
	}
	if user != nil {//用户名重复
		t := template.Must(template.ParseFiles("views/pages/user/regist.html"))
		t.Execute(w,"用户名已存在！请重新输入。")
	}else{
		//插入用户
		errIns := dao.AddUser(username,password,email)
		if errIns != nil {
			fmt.Println("DoRegister error:",errIns)
		}
		t := template.Must(template.ParseFiles("views/pages/user/regist_success.html"))
		t.Execute(w,"")
	}
}

//检查用户名是否存在
func DoFindUserByName(w http.ResponseWriter,r *http.Request){
	//获取参数
	username := r.PostFormValue("username")

	//查询是否存在用户
	user,err := dao.CheckUserName(username)
	if err != nil {
		fmt.Println("DoFindUserByName error:",err)
	}
	if user != nil { //用户名重复
		w.Write([]byte("用户名已存在"))
	}else{
		w.Write([]byte("<font style='color:green'>恭喜,用户名可用</font>"))
	}
}

//处理注销请求
func LogoutHandler(w http.ResponseWriter,r *http.Request){
	//获取要注销的会话
	session,err := dao.GetSessionByCookie(r)
	if err != nil{
		fmt.Println("LogoutHandler error:",err)
	}else{
		//删除会话
		errDel := dao.DeleteSession(session.SessionId)
		if errDel != nil{
			fmt.Println("LogoutHandler-DeleteSession error:",errDel)
		}
	}

	//设置cookie失效
	cookie,errCookie := r.Cookie("sessionId")
	if errCookie != nil{
		fmt.Println("LogoutHandler error:",errCookie)
	}else{
		cookie.MaxAge = -1
		http.SetCookie(w,cookie)
	}

	//退出到首页
	IndexHandler(w,r)
}