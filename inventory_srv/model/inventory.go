package model


type Inventory struct {
	BaseModel

	Goods   int32 `gorm:"type:int;index"`  // 商品 id
	Stocks  int32 `gorm:"type:int"`        // 库存，涉及仓库就复杂了
	Version int32 `gorm:"type:int"`        // 分布式锁的乐观锁
}

func (i Inventory) TableName() string {
	return "inventory"
}


// 订单历史
type InventoryHistory struct {
	User   int32
	Goods  int32
	Num    int32
	Order  int32
	Status int32  // 1. 表示库存是预扣减，幂等性；2 表示已经支付
}
