package model




import (
	"fmt"
	"time"

	"gorm.io/gorm"
)

type BaseModel struct {
	ID        uint64         `gorm:"column:id;primaryKey;autoIncrement" json:"id,omitempty"` 
	CreatedAt time.Time      `gorm:"column:created_at;index" json:"created_at,omitempty"`
	UpdatedAt time.Time      `gorm:"column:updated_at;index" json:"updated_at,omitempty"`
	DeletedAt gorm.DeletedAt 
}


func (a BaseModel) GeyStringID() string {
	return fmt.Sprintf("%v", a.ID)
} 

/*
 
模板，就是小红岛

 */