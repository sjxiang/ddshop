package dal

import "github.com/sjxiang/ddshop/user_srv/cmd/user/dal/db"


// Init 初始化数据库
func Init() {
	db.Init()  // mysql
}