package handler

import (
	"GopherMall/inventory_srv/global"
	"GopherMall/inventory_srv/model"
	proto "GopherMall/inventory_srv/proto/.InventoryProto"
	"GopherMall/inventory_srv/utils"
	"context"
	"errors"
	"fmt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
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

func (i InventoryServer) Sell(ctx context.Context, req *proto.SellInfo) (*proto.Empty, error) {

	mutex := utils.RedisSync.NewMutex("goods")

	if err := mutex.Lock(); err != nil {
		return nil, status.Errorf(codes.Internal, "创建Redis互斥锁失败")
	}

	tx := global.DB.Begin()

	for _, goodInfo := range req.GoodsInfo {

		var inv model.Inventory

		if result := global.DB.Clauses(clause.Locking{Strength: "UPDATE"}).Where(&model.Inventory{Goods: goodInfo.GoodsId}).First(&inv); result.Error != nil {
			if errors.Is(result.Error, gorm.ErrRecordNotFound) {
				tx.Rollback()
				return nil, status.Errorf(codes.NotFound, "无效ID")
			}
			tx.Rollback()
			return nil, status.Errorf(codes.Internal, "服务器错误")
		}

		if inv.Stocks < goodInfo.Num {
			tx.Rollback()
			return nil, status.Errorf(codes.InvalidArgument, "库存不足")
		}

		inv.Stocks -= goodInfo.Num
		tx.Save(&inv)

	}
	tx.Commit()

	_, _ = mutex.Unlock()

	return &proto.Empty{}, nil
}

func (i InventoryServer) Reback(ctx context.Context, req *proto.SellInfo) (*proto.Empty, error) {

	tx := global.DB.Begin()
	for _, goodInfo := range req.GoodsInfo {

		mutex := utils.RedisSync.NewMutex(fmt.Sprintf("goods_%d", goodInfo.GoodsId))
		if err := mutex.Lock(); err != nil {
			tx.Rollback()
			return nil, status.Errorf(codes.Internal, "创建Redis互斥锁失败")
		}

		var inv model.Inventory
		if result := global.DB.Clauses(clause.Locking{Strength: "UPDATE"}).Where(&model.Inventory{Goods: goodInfo.GoodsId}).Find(&inv); result.RowsAffected == 0 {
			tx.Rollback()
			return nil, status.Errorf(codes.NotFound, "无效ID")
		}
		inv.Stocks = inv.Stocks + goodInfo.Num
		tx.Save(&inv)

		if ok, err := mutex.Unlock(); !ok || err != nil {
			tx.Rollback()
			return nil, status.Errorf(codes.Internal, "解除Redis互斥锁失败")
		}
	}
	tx.Commit()

	return &proto.Empty{}, nil
}
