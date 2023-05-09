package service

import (
	"time"

	"golang.org/x/crypto/bcrypt"
)

// 获取页码信息
func GetPageOffset(pageNum, pageSize uint32) int {
	if pageNum < 1 {
		pageNum = 1
	}
	return int((pageNum - 1) * pageSize)
}

// 格式化时间
func DateFormat(times int64) string {
	return time.Unix(times, 0).Format("2006-01-02")
}


func hashPassowrd(pwd string) string {
	bytes, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}
	return string(bytes)
}

func verifyPassowrd(hashed, plainText string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashed), []byte(plainText))
	return err == nil 
}