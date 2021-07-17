package dao

import (
	"testing"
)

func TestGetBooks(t *testing.T) {
	res,err := GetBooks()
	if err != nil {
		t.Error(err)
	}
	for _,book := range res {
		t.Log("TestGetBook-book:",book)
	}
}