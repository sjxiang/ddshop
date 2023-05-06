package model

import (
	"errors"
	"regexp"
	"time"

	"golang.org/x/crypto/bcrypt"
)


type User struct {
	BaseModel

	Mobile   string     `gorm:"index:idx_mobile;unique;type:varchar(11);not null"`
	Email    string     `gorm:"index:idx_email;unique;type:varchar(36);not null"`
	NickName string     `gorm:"type:varchar(20)"`  // 系统自动生成，用户可以自己修改
	Password string     `gorm:"type:varchar(100);not null"`
	Birthday *time.Time `gorm:"type:datetime;default:2023-05-01 01:02:03.000"`
	Avatar   string     `gorm:"type:varchar(60)"`  
	Gender   string     `gorm:"column:gender;default:male;type:varchar(6) comment 'female 表示女生，male 表示男生'"`
	Role     int        `gorm:"column:role;default:1;type:int comment '1 表示普通用户，2 表示管理员'"`
}

func (u User) TableName() string {
	return "user"
}


func (u User) Check() error {
	const emailPattern = "(.+)@(.+){2,}\\.(.+){2,}"

	// if len(u.Password) < 8 {
	// 	return errors.New("密码长度过短")
	// }

	if len(u.Mobile) != 11 {
		return errors.New("手机号长度必须为 11 位")
	}

	ok, err := regexp.Match(emailPattern, []byte(u.Email))
	if !ok || err !=nil {
		return errors.New("不正确的邮箱地址")
	}
	
	return nil
}


/*

信息摘要算法

	e.g. v2ex 7 年前端求职，无能狂怒，破防

暴力破解，彩虹表遍历

*/


func hashPassowrd(pwd string) string {
	bytes, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}
	return string(bytes)
}

func verifyPassowrd(hashed, input string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashed), []byte(input))
	return err == nil 
}
