package db

import (
	"context"
	"time"

	"github.com/sjxiang/ddshop/user_srv/model"
	"gorm.io/gorm"
)

// 获取用户信息列表
func MGetUser(ctx context.Context, limit, offset int) ([]*model.User, int64, error) {

	var total int64
	var res []*model.User
	conn := DB.WithContext(ctx).Model(new(model.User))

	if err := conn.Count(&total).Error; err != nil {
		return res, total, err
	}

	if err := conn.Limit(limit).Offset(offset).Find(&res).Error; err != nil {
		return res, total, err
	}

	return res, total, nil
}

func QueryUserByID(ctx context.Context, userIDs ...int32) ([]*model.User, error) {
	var res []*model.User
	if len(userIDs) == 0 {
		return res, nil 
	}

	if err := DB.WithContext(ctx).Where("id in ?", userIDs).Find(&res).Error; err != nil {
		return res, err
	}
	
	return res, nil
}

func QueryUserByEmail(ctx context.Context, email string) (*model.User, error) {
	var res model.User

	if err := DB.WithContext(ctx).Where("email = ?", email).Find(&res).Error; err != nil {
		return &res, err
	}
	
	// 查询错误和数据不存在，没有区分 

	return &res, nil
}

func QueryUserByEmailPlus(ctx context.Context, email string, user *model.User) *gorm.DB {
	return DB.WithContext(ctx).Where("email = ?", email).Find(user)
}

func CreateUser(ctx context.Context, user *model.User) error {
	if err := DB.WithContext(ctx).Model(new(model.User)).Create(user).Error; err != nil {
		return err
	}
	return nil 
}

func UpdateUser(ctx context.Context, userID int32, mobile, email, nickname *string, password, birthday *int64) error {
	params := map[string]interface{}{}
	if mobile != nil {
		params["mobile"] = *mobile
	}
	if email != nil {
		params["content"] = *email
	}
	if nickname != nil {
		params["nickname"] = *nickname
	}
	if password != nil {
		params["password"] = *password  // 加密，在 db 层搞这个，多余且不好看
	}
	if birthday != nil {
		params["birthday"] = time.Unix(*birthday, 0)  // 时间戳转换
	}

	return DB.WithContext(ctx).Model(new(model.User)).Where("id = ?", userID).
		Updates(params).Error

	// 或者 save()
}
