package pack

import (
	"github.com/sjxiang/ddshop/inventory_srv/model"
	"github.com/sjxiang/ddshop/inventory_srv/pb"
)

// "github.com/sjxiang/ddshop/user_srv/model"
// "github.com/sjxiang/ddshop/user_srv/pb"

// 序列化

func InventoryInfo(m *model.Inventory) *pb.GoodsInvInfo {
	if m == nil {
		return nil
	}

	// tips：在 gRPC 的 message 中字段有默认值，不能随便复制 nil 进去，容易出错

	return &pb.GoodsInvInfo{
		GoodsId: m.Goods,
		Num: m.Stocks,
	}
}

func InvInfoList(ms []*model.Inventory) []*pb.GoodsInvInfo {
	invList := make([]*pb.GoodsInvInfo, 0)
	for _, m := range ms {
		if u := InventoryInfo(m); u != nil {
			invList = append(invList, u)
		}
	}
	return invList
}

