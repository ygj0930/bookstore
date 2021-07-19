package daotest

import (
	"bookstore/dao"
	"bookstore/model"
	"testing"
)

func TestAddSession(t *testing.T){
	session := &model.Session{
		SessionId: "session_1",
		UserName:  "user_1",
		UserId:   1,
	}

	err := dao.AddSession(session)

	if err != nil {
		t.Log(err)
	}
}

func TestGetSessionByID(t *testing.T){
	session,err := dao.GetSessionByID("session_1")

	if err != nil {
		t.Log(err)
	}else{
		t.Logf("TestGetSessionByID:%+v",session)
	}
}

func TestDeleteSession(t *testing.T){
	err := dao.DeleteSession("session_1")

	if err != nil {
		t.Log(err)
	}
}