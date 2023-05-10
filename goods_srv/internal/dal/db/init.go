package db

import (
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"github.com/sjxiang/ddshop/goods_srv/model"
)

var DB *gorm.DB

// Init init DB
func Init() {
	var err error
	
	dsn := os.Getenv("DSN")

	DB, err = gorm.Open(mysql.Open(dsn),
		&gorm.Config{
			SkipDefaultTransaction: true,  // 关闭默认事务
			PrepareStmt: true,             // 缓存预编译语句
			// DisableForeignKeyConstraintWhenMigrating: true,  // 不允许外键
			Logger: logger.Default.LogMode(logger.Info),
		},
	)
	if err != nil {
		panic("failed to connect database" + err.Error())
	}

	if err := DB.AutoMigrate(new(model.Goods), new(model.Banner), new(model.Brands), new(model.Category), new(model.GoodsCategoryBrand)); err != nil {
		panic(err)
	} 
		
}

