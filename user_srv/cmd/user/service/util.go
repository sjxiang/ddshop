package service

import "golang.org/x/crypto/bcrypt"


func GetPageOffset(pageNum, pageSize uint32) int {
	return int((pageNum - 1) * pageSize)
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