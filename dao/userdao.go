package dao

import (
	"bookstore/utils"
	"bookstore/model"
	"fmt"
)

//CheckUserNameAndPassword 验证用户名与密码，查询对应用户记录
func CheckUserNameAndPassword(username string,password string)(user *model.User,err error){
	sqlStr := "select * from users where username = ? and password = ?"

	row := utils.Db.QueryRow(sqlStr,username,password)

	//提取sql结果
	user = &model.User{}
	errScan :=row.Scan(&user.ID,&user.Username,&user.Password,&user.Email)
	if errScan!=nil {
		fmt.Println("CheckUserNameAndPassword Scan err:",errScan)
		return nil,errScan
	}
	return user,nil
}

//CheckUserName 验证用户是否存在
func CheckUserName(username string)(user *model.User,err error){
	sqlStr := "select * from users where username = ?"

	row := utils.Db.QueryRow(sqlStr,username)

	//提取sql结果
	user = &model.User{}
	errScan :=row.Scan(&user.ID,&user.Username,&user.Password,&user.Email)
	if errScan!=nil {
		fmt.Println("CheckUserName Scan err:",errScan)
		return nil,errScan
	}
	return user,nil
}

//AddUser 新增用户
func AddUser(username string,password string,email string) error {
	sqlStr := "insert into users(username,password,email) values(?,?,?)"

	_,errIns := utils.Db.Exec(sqlStr,username,password,email)
	if errIns != nil {
		return errIns
	}
	return nil
}