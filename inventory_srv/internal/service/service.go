package service

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	emptypb "google.golang.org/protobuf/types/known/emptypb"

	"github.com/sjxiang/ddshop/inventory_srv/internal/dal/db"
	"github.com/sjxiang/ddshop/inventory_srv/pb"
	"github.com/sjxiang/ddshop/inventory_srv/internal/pack"
)

type InventoryServiceImpl struct {
	pb.UnimplementedInventoryServer
}


func (InventoryServiceImpl) SetInv(ctx context.Context, req *pb.GoodsInvInfo) (*emptypb.Empty, error) {
	
	// 先看看有没有
	_, err := db.MGetInventory(ctx, req.GoodsId)
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}
	
	// 更新
	if err := db.UpdateInventory(ctx, req.GoodsId, req.Num); err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}
	
	return &emptypb.Empty{}, nil
}

func (InventoryServiceImpl) InvDetail(ctx context.Context, req *pb.GoodsInvInfo) (*pb.GoodsInvInfo, error) {
	res, err := db.MGetInventory(ctx, req.GoodsId)
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}
	
	return pack.InvInfoList(res)[0], nil
}

func (InventoryServiceImpl) Sell(ctx context.Context, req *pb.SellInfo) (*emptypb.Empty, error) {
	

	// 包装过的 dao 没法整事务
	
	for _, goodInfo := range req.Data {
	
		// 先看看有没有
		res, err := db.MGetInventory(ctx, goodInfo.GoodsId)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "没有库存信息")
		}
		
		// 判断库存是否充足
		balance := res[0].Stocks - goodInfo.Num 
		if balance < 0 {
			return nil, status.Errorf(codes.Internal, "库存不足")
		}

		// 更新
		if err := db.UpdateInventory(ctx, goodInfo.GoodsId, balance); err != nil {
			return nil, status.Errorf(codes.Internal, err.Error())
		}
		
	}

	return &emptypb.Empty{}, nil
}

func (InventoryServiceImpl) Reback(ctx context.Context, req *pb.SellInfo) (*emptypb.Empty, error) {

	// 库存归还
	// 1. 订单超时归还，30 min 没付钱
	// 2. 订单创建失败
	// 3. 用户手动取消





	return nil, nil
}