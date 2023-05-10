package model



type GoodsCategoryBrand struct {
	BaseModel

	CategoryID uint64 `gorm:"index:idx_category_brands,unique"`
	Category   Category

	BrandsID   uint64 `gorm:"index:idx_category_brands,unique"`
	Brands     Brands
}


func (g GoodsCategoryBrand) TableName() string {
	return "goods_category_brand"
}
