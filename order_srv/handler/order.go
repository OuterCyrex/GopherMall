package handler

import (
	goodsproto "GopherMall/goods_srv/proto/.GoodsProto"
	invproto "GopherMall/inventory_srv/proto/.InventoryProto"
	"GopherMall/order_srv/global"
	"GopherMall/order_srv/model"
	proto "GopherMall/order_srv/proto/.OrderProto"
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type OrderServer struct {
	proto.UnimplementedOrderServer
}

func (o OrderServer) CreateOrder(ctx context.Context, req *proto.OrderRequest) (*proto.OrderInfoResponse, error) {
	var goodsIds []int32
	var shopCarts []model.ShoppingCart
	goodsNumMap := make(map[int32]int32)

	result := global.DB.Where(&model.ShoppingCart{User: req.UserId, Checked: true}).Find(&shopCarts)
	if result.RowsAffected == 0 {
		return nil, status.Errorf(codes.NotFound, "订单不存在")
	}

	for _, shop := range shopCarts {
		goodsIds = append(goodsIds, shop.Goods)
		goodsNumMap[shop.Goods] = shop.Nums
	}

	goodsResp, err := global.GoodsSrvClient.BatchGetGoods(ctx, &goodsproto.BatchGoodsIdInfo{Id: goodsIds})
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	var amount float32
	var orderGoods []*model.OrderGoods
	var sellList []*invproto.GoodsInvInfo
	for _, goods := range goodsResp.Data {
		amount += goods.ShopPrice * float32(goodsNumMap[goods.Id])
		orderGoods = append(orderGoods, &model.OrderGoods{
			Goods:      goods.GetId(),
			GoodsName:  goods.Name,
			GoodsImage: goods.GoodsFrontImage,
			GoodsPrice: goods.MarketPrice,
			Num:        goodsNumMap[goods.Id],
		})

		sellList = append(sellList, &invproto.GoodsInvInfo{
			GoodsId: goods.Id,
			Num:     goodsNumMap[goods.Id],
		})
	}

	_, err = global.InventorySrvClient.Sell(ctx, &invproto.SellInfo{GoodsInfo: sellList})
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	// TODO

	return nil, nil
}

func (o OrderServer) OrderList(ctx context.Context, req *proto.OrderFilterRequest) (*proto.OrderListResponse, error) {
	var orders []model.OrderInfo
	var count int64

	global.DB.Where(&model.OrderInfo{User: req.UserId}).Count(&count)

	result := global.DB.Scopes(Paginate(int(req.Pages), int(req.PagePerNums))).Where(&model.OrderInfo{User: req.UserId}).Find(&orders)
	if result.Error != nil {
		return nil, result.Error
	}

	var respList []*proto.OrderInfoResponse

	for _, order := range orders {
		respList = append(respList, &proto.OrderInfoResponse{
			Id:      order.ID,
			UserId:  order.User,
			OrderSn: order.OrderSn,
			PayType: order.PayType,
			Status:  order.Status,
			Post:    order.Post,
			Total:   order.OrderMount,
			Address: order.Address,
			Name:    order.SignerName,
			Mobile:  order.SignerMobile,
		})
	}

	return &proto.OrderListResponse{
		Total: int32(count),
		Data:  respList,
	}, nil
}

func (o OrderServer) OrderDetail(ctx context.Context, req *proto.OrderRequest) (*proto.OrderInfoDetailResponse, error) {
	var order model.OrderInfo

	result := global.DB.Where(&model.OrderInfo{BaseModel: model.BaseModel{ID: req.Id}}).Find(&order)
	if result.RowsAffected == 0 {
		return nil, status.Errorf(codes.NotFound, "无效ID")
	}

	var orderGoods []model.OrderGoods
	var respList []*proto.OrderItemResponse
	result = global.DB.Where(&model.OrderGoods{Order: order.ID}).Find(&orderGoods)
	if result.Error != nil {
		return nil, result.Error
	}

	for _, good := range orderGoods {
		respList = append(respList, &proto.OrderItemResponse{
			Id:         good.ID,
			OrderId:    good.Order,
			GoodsId:    good.Goods,
			GoodsName:  good.GoodsName,
			GoodsImage: good.GoodsImage,
			GoodsPrice: good.GoodsPrice,
			Nums:       good.Num,
		})
	}

	return &proto.OrderInfoDetailResponse{
		OrderInfo: &proto.OrderInfoResponse{
			Id:      order.ID,
			UserId:  order.User,
			OrderSn: order.OrderSn,
			PayType: order.PayType,
			Status:  order.Status,
			Post:    order.Post,
			Total:   order.OrderMount,
			Address: order.Address,
			Name:    order.SignerName,
			Mobile:  order.SignerMobile,
		},
		Goods: respList,
	}, nil
}

func (o OrderServer) UpdateOrderStatus(ctx context.Context, req *proto.OrderStatus) (*proto.Empty, error) {
	return nil, nil
}
