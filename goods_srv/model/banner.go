package model


// 轮播图
type Banner struct {
	BaseModel

	Image string `gorm:"type:varchar(200);not null"`
	Url   string `gorm:"type:varchar(200);not null"`
	Index int32  `gorm:"type:int;default 1;not null"`  // 顺序
}

func (b Banner) TableName() string {
	return "banner"
}
