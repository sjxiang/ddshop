package model

import "time"


type ShoppingCart struct {
	BaseModel

	User    int32 `gorm:"type:int;index"`  // 方便查询当前用户的购物车记录
	Goods   int32 `gorm:"type:int;index"`
	Nums    int32 `gorm:"type:int;index"`
	Checked bool                           // 是否选中 
}

func (ShoppingCart) TableName() string {
	return "shopping_cart"
}


type OrderInfo struct {
	BaseModel

	User         int32      `gorm:"type:int;index"`
	OrderSn      string     `gorm:"type:varchar(30);index"`  // 订单号，自己平台生成的订单号
	PayType      string     `gorm:"type:varchar(20);comment 'alipay(支付宝) wechat(微信)'"`

	Status       string     `gorm:"type:varchar(20);comment 'PAYING(待支付) TRADE_SUCCESS(成功) TRADE_CLOSED(超时关闭) WAIT_BUYER_PAY(交易创建) TRADE_FINISH(交易结束)'"`
	TradeNo      string     `gorm:"type:varchar(100);comment '交易号'"`  // 支付宝的订单号，便于查账
	OrderMount   float32    // 订单金额
	PayTime      time.Time

	Address      string     `gorm:"type:varchar(100)"` 
	SignerName   string     `gorm:"type:varchar(20)"` 
	SignerMobile string     `gorm:"type:varchar(11)"` 
	Post         string     `gorm:"type:varchar(20)"` 
}

func (OrderInfo) TableName() string {
	return "order_info"
}


type OrderGoods struct {
	BaseModel

	Order int32 `gorm:"type:int;index"`
	Goods int32 `gorm:"type:int;index"`
	

	// 
	GoodsName  string `gorm:"type:varchar(100);index"` 
	GoodsImage string `gorm:"type:varchar(200)"` 
	GoodsPrice string  
	Nums       int32  `gorm:"type:int"`
}

func (OrderGoods) TableName() string {
	return "order_goods"
}