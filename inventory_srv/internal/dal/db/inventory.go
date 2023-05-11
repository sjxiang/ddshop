package db

import (
	"context"

	"github.com/sjxiang/ddshop/inventory_srv/model"
)



// // CreateNote create note info
// func CreateNote(ctx context.Context, notes []*Note) error {
// 	if err := DB.WithContext(ctx).Create(notes).Error; err != nil {
// 		return err
// 	}
// 	return nil
// }

// MGetNotes multiple get list of note info
func MGetInventory(ctx context.Context, inventoryIDs ...int32) ([]*model.Inventory, error) {
	var res []*model.Inventory
	if len(inventoryIDs) == 0 {
		return res, nil
	}

	if err := DB.WithContext(ctx).Where("goods in ?", inventoryIDs).Find(&res).Error; err != nil {
		return res, err
	}

	return res, nil
}

// // UpdateNote update note info
func UpdateInventory(ctx context.Context, goodsID,  num int32) error {
	params := map[string]interface{}{}
	params["stocks"] = num
	return DB.WithContext(ctx).Model(&model.Inventory{}).Where("goods = ?", goodsID).
		Updates(params).Error
}

// // DeleteNote delete note info
// func DeleteNote(ctx context.Context, noteID, userID int64) error {
// 	return DB.WithContext(ctx).Where("id = ? and user_id = ? ", noteID, userID).Delete(&Note{}).Error
// }

// // QueryNote query list of note info
// func QueryNote(ctx context.Context, userID int64, searchKey *string, limit, offset int) ([]*Note, int64, error) {
// 	var total int64
// 	var res []*Note
// 	conn := DB.WithContext(ctx).Model(&Note{}).Where("user_id = ?", userID)

// 	if searchKey != nil {
// 		conn = conn.Where("title like ?", "%"+*searchKey+"%")
// 	}

// 	if err := conn.Count(&total).Error; err != nil {
// 		return res, total, err
// 	}

// 	if err := conn.Limit(limit).Offset(offset).Find(&res).Error; err != nil {
// 		return res, total, err
// 	}

// 	return res, total, nil
// }
