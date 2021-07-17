package dao

import (
	"testing"
)

func TestCheckUserName(t *testing.T) {
	user,err := CheckUserName("lili")
	if err!=nil {
		t.Error(err)
	}
	t.Log("TestCheckUserName:",user)
}

func TestCheckUserNameAndPassword(t *testing.T) {
	user,err := CheckUserNameAndPassword("lili","e10adc3949ba59abbe56e057f20f883e")
	if err!=nil {
		t.Error(err)
	}
	t.Log("TestCheckUserNameAndPassword:",user)
}

func TestAddUser(t *testing.T) {
	err := AddUser("admin1","123456","admin@123.com")
	if err!=nil {
		t.Error(err)
	}
	t.Log("TestAddUser success")
}