package pack

import (
	"github.com/sjxiang/ddshop/user_srv/model"
	"github.com/sjxiang/ddshop/user_srv/pb"
)

// 序列化

func UserInfo(m *model.User) *pb.UserInfo {
	if m == nil {
		return nil
	}

	// tips：在 gRPC 的 message 中字段有默认值，不能随便复制 nil 进去，容易出错

	return &pb.UserInfo{
		Id:       m.ID,
		Mobile:   m.Mobile,
		Email:    m.Email,
		Password: m.Password,
		Nickname: m.NickName,
		Birthday: uint64(m.Birthday.Unix()),
		// Avatar:   m.Avatar,
		Role:     int32(m.Role),
	}
}

func UserInfoList(ms []*model.User) []*pb.UserInfo {
	userList := make([]*pb.UserInfo, 0)
	for _, m := range ms {
		if u := UserInfo(m); u != nil {
			userList = append(userList, u)
		}
	}
	return userList
}

