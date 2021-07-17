package controller

import (
	"bookstore/dao"
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
		t := template.Must(template.ParseFiles("views/pages/user/login_success.html"))
		t.Execute(w,username)
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