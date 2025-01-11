package handler

import (
	"GopherMall/order_srv/global"
	"GopherMall/order_srv/model"
	proto "GopherMall/order_srv/proto/.OrderProto"
	"context"
	"errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
)

func (o OrderServer) CartItemList(ctx context.Context, req *proto.UserInfo) (*proto.CartItemListResponse, error) {
	var shopCarts []model.ShoppingCart

	result := global.DB.Where(&model.ShoppingCart{User: req.Id}).Find(&shopCarts)
	if result.Error != nil {
		return nil, result.Error
	}

	var respList []*proto.ShopCartInfoResponse

	for _, shopCart := range shopCarts {
		respList = append(respList, &proto.ShopCartInfoResponse{
			Id:      shopCart.ID,
			UserId:  shopCart.User,
			GoodsId: shopCart.Goods,
			Nums:    shopCart.Nums,
			Checked: shopCart.Checked,
		})
	}

	return &proto.CartItemListResponse{
		Total: int32(result.RowsAffected),
		Data:  respList,
	}, nil
}
func (o OrderServer) CreateCartItem(ctx context.Context, req *proto.CartItemRequest) (*proto.ShopCartInfoResponse, error) {
	var shopCart model.ShoppingCart

	result := global.DB.Where(&model.ShoppingCart{User: req.UserId, Goods: req.GoodsId}).First(&shopCart)
	if result.Error != nil && !errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, result.Error
	}

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		shopCart.Nums += req.Nums
	} else {
		shopCart = model.ShoppingCart{
			User:    req.UserId,
			Goods:   req.GoodsId,
			Nums:    req.Nums,
			Checked: false,
		}
	}

	err := global.DB.Save(&shopCart).Error
	if err != nil {
		return nil, err
	}

	return &proto.ShopCartInfoResponse{
		Id: shopCart.ID,
	}, nil
}

func (o OrderServer) UpdateCartItem(ctx context.Context, req *proto.CartItemRequest) (*proto.Empty, error) {
	var shopCart model.ShoppingCart

	result := global.DB.Where("Id = ?", req.Id).Find(&shopCart)
	if result.Error != nil {
		return nil, result.Error
	}

	if result.RowsAffected == 0 {
		return nil, status.Errorf(codes.NotFound, "无效ID")
	}

	var updateMap map[string]interface{}

	if shopCart.Checked != req.Checked {
		updateMap["checked"] = req.Checked
	}

	err := global.DB.Updates(model.ShoppingCart{
		User:  req.UserId,
		Goods: req.GoodsId,
		Nums:  req.Nums,
	}).Updates(updateMap).Where("Id = ?", req.Id).Error
	if err != nil {
		return nil, err
	}

	return &proto.Empty{}, nil
}

func (o OrderServer) DeleteCartItem(ctx context.Context, req *proto.CartItemRequest) (*proto.Empty, error) {
	result := global.DB.Delete(&model.ShoppingCart{}, req.Id)
	if result.RowsAffected == 0 {
		return nil, status.Errorf(codes.NotFound, "无效ID")
	}

	if result.Error != nil {
		return nil, result.Error
	}

	return &proto.Empty{}, nil
}
