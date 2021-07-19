package dao

import (
	"bookstore/model"
	"bookstore/utils"
	"net/http"
)

//保存会话
func AddSession(s *model.Session) error{
	sqlStr := "insert into sessions values(?,?,?)"

	_,errIns :=utils.Db.Exec(sqlStr,s.SessionId,s.UserName,s.UserId)
	if errIns != nil {
		return errIns
	}

	return nil
}


//删除会话
func DeleteSession(sessionId string) error {
	sqlStr := "delete from sessions where session_id = ?"

	_,errDel :=utils.Db.Exec(sqlStr,sessionId)

	if errDel != nil {
		return errDel
	}

	return nil
}

//查询某个会话
func GetSessionByID(sessionId string) (*model.Session,error){
	sqlStr := "select * from sessions where session_id = ?"
	row := utils.Db.QueryRow(sqlStr,sessionId)

	session := &model.Session{}

	errScan := row.Scan(&session.SessionId,&session.UserName,&session.UserId)
	if errScan != nil {
		return nil,errScan
	}
	return session,nil
}

//从请求头提取会话id，并查询会话返回
func GetSessionByCookie(r *http.Request)(*model.Session,error){
	cookie,err := r.Cookie("sessionId")
	if err != nil{
		return nil, err
	}
	session,errGet := GetSessionByID(cookie.Value)
	if errGet != nil {
		return nil, err
	}
	return session,nil
}

//判断是否登录
func IsLogin(r *http.Request) bool {
	session,err := GetSessionByCookie(r)

	if err != nil || session == nil {
		return false
	}
	return true
}