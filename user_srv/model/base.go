package model


import (
	"time"

	"gorm.io/gorm"
)

type BaseModel struct {
	ID        int32      `gorm:"primarykey"`
	CreatedAt time.Time  
	UpdatedAt time.Time  
	DeletedAt gorm.DeletedAt
	IsDeleted bool       
}
