package model


// 商品
type Goods struct {
	BaseModel

	CategoryID  uint64  `gorm:"not null"`
	Category    Category
	BrandsID    uint64  `gorm:"not null"`
	Brands      Brands

	Onsale      bool    `gorm:"default:false;not null"`  // 是否上架
	ShipFree    bool    `gorm:"default:false;not null"`  // 免运费
	IsNew       bool    `gorm:"default:false;not null"`  // 新品
	IsHot       bool    `gorm:"default:false;not null"`  // 爆款 

	Name        string  `gorm:"type:varchar(50);not null"`  
	Sn          string  `gorm:"type:varchar(50);not null"` // 商品编码
	ClickNum    uint64  `gorm:"default:0;not null"`  // 点击量 
 	SoldNum     uint64  `gorm:"default:0;not null"`  // 销量
	FavNum      uint64  `gorm:"default:0;not null"`  // 收藏数
	MarketPrice float32 `gorm:"not null"`     // 市场价
	ShopPrice   float32 `gorm:"not null"`  // 会员价
	Intro       string  `gorm:"type:varchar(100);not null"` // 商品简介
	FrontImage  string  `gorm:"type:varchar(200);not null"` // 封面图
}

