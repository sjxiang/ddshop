package model


/*

上架的时候，需要投稿


参数
	品牌 - Cisco、Huawei
	类目 - 交换机、路由器、核心堆叠、三层、二层傻瓜交换机

 */


// 分类
type Category struct {
	BaseModel

	ParentCategoryID uint64
	ParentCategory   *Category
	
	Name  string `gorm:"type:varchar(20);not null"`
	Level int32  `gorm:"type:int;not null;default:1"`
	IsTab bool   `gorm:"default:false;not null"`
}

func (c Category) TableName() string {
	return "category"
}


