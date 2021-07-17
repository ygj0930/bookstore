package utils

import (
	"crypto/md5"
	"fmt"
)

//Md5加密
func Md5(unEncry string) (encrytion string){
	byte16 := []byte(unEncry)
	encrytion = fmt.Sprintf("%x",md5.Sum(byte16))
	return
}
