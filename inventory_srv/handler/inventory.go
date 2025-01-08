package handler

import (
	"GopherMall/inventory_srv/global"
	"GopherMall/inventory_srv/model"
	proto "GopherMall/inventory_srv/proto/.InventoryProto"
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm/clause"
	"sync"
)

type InventoryServer struct {
	proto.UnimplementedInventoryServer
}

func (i InventoryServer) SetInv(ctx context.Context, req *proto.GoodsInvInfo) (*proto.Empty, error) {
	inv := model.Inventory{}
	_ = global.DB.Where("goods = ?", req.GoodsId).Find(&inv)

	inv.Goods = req.GetGoodsId()
	inv.Stocks = req.GetNum()

	err := global.DB.Save(&inv).Error
	if err != nil {
		return nil, status.Errorf(codes.Internal, "数据库出错")
	}

	return &proto.Empty{}, nil
}

func (i InventoryServer) InvDetail(ctx context.Context, req *proto.GoodsInvInfo) (*proto.GoodsInvInfo, error) {
	inv := model.Inventory{}
	result := global.DB.Where("goods = ?", req.GoodsId).Find(&inv)

	if result.RowsAffected == 0 {
		return nil, status.Errorf(codes.NotFound, "无效ID")
	}

	return &proto.GoodsInvInfo{
		GoodsId: inv.Goods,
		Num:     inv.Stocks,
	}, nil
}

var m sync.Mutex

// Sell 使用了悲观锁
func (i InventoryServer) Sell(ctx context.Context, req *proto.SellInfo) (*proto.Empty, error) {
	// 上锁
	m.Lock()

	tx := global.DB.Begin()
	for _, goodInfo := range req.GoodsInfo {
		var inv model.Inventory
		if result := global.DB.Clauses(clause.Locking{Strength: "UPDATE"}).Where(&model.Inventory{Goods: goodInfo.GoodsId}).Find(&inv); result.RowsAffected == 0 {
			tx.Rollback()
			return nil, status.Errorf(codes.NotFound, "无效ID")
		}
		if inv.Stocks < goodInfo.Num {
			tx.Rollback()
			return nil, status.Errorf(codes.InvalidArgument, "库存不足")
		}
		inv.Stocks -= goodInfo.Num
		tx.Save(&inv)
	}
	tx.Commit()

	// 上锁
	m.Unlock()

	return &proto.Empty{}, nil
}

// Reback 使用了乐观锁
func (i InventoryServer) Reback(ctx context.Context, req *proto.SellInfo) (*proto.Empty, error) {

	tx := global.DB.Begin()
	for _, goodInfo := range req.GoodsInfo {
		var inv model.Inventory
		for {
			if result := global.DB.Where(&model.Inventory{Goods: goodInfo.GoodsId}).Find(&inv); result.RowsAffected == 0 {
				tx.Rollback()
				return nil, status.Errorf(codes.NotFound, "无效ID")
			}
			if result := tx.Model(&model.Inventory{}).Select("stocks", "version").Where("goods = ? AND version = ?", inv.Goods, inv.Version).Updates(&model.Inventory{
				Stocks:  inv.Stocks + goodInfo.Num,
				Version: inv.Version + 1,
			}); result.RowsAffected != 0 {
				break
			}
		}
	}
	tx.Commit()

	return &proto.Empty{}, nil
}
