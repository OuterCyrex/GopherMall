package handler

import (
	"GopherMall/goods_srv/global"
	"GopherMall/goods_srv/model"
	proto "GopherMall/goods_srv/proto/.GoodsProto"
	"context"
	"fmt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type GoodsServer struct {
	proto.UnimplementedGoodsServer
}

func (g GoodsServer) GoodsList(ctx context.Context, req *proto.GoodsFilterRequest) (*proto.GoodsListResponse, error) {
	if req.PriceMin > req.PriceMax {
		return nil, status.Errorf(codes.InvalidArgument, "最低价格大于最高价格")
	}

	tx := global.DB.Model(&model.Goods{})

	if req.KeyWords != "" {
		tx = tx.Where("name LIKE ?", "%"+req.KeyWords+"%")
	}

	if req.IsHot {
		tx = tx.Where("is_hot = ?", req.IsHot)
	}

	if req.IsNew {
		tx = tx.Where("is_new = ?", req.IsNew)
	}

	if req.IsTab {
		tx = tx.Where("is_tab = ?", req.IsTab)
	}

	if req.PriceMin > 0 {
		tx = tx.Where("shop_price >= ?", req.PriceMin)
	}

	if req.PriceMax > 0 {
		tx = tx.Where("shop_price <= ?", req.PriceMax)
	}

	if req.Brand > 0 {
		tx = tx.Where("brands_id = ?", req.Brand)
	}

	if req.TopCategory > 0 {
		var category model.Category
		subQuery := ""

		result := global.DB.Model(model.Category{}).Where("Id = ?", req.TopCategory).Find(&category)
		if result.RowsAffected == 0 {
			return nil, status.Errorf(codes.NotFound, "无效标签ID")
		}

		switch category.Level {
		case 1:
			subQuery = fmt.Sprintf("SELECT id FROM category WHERE parent_category_id IN (SELECT id FROM category WHERE parent_category_id = %d)", req.TopCategory)
		case 2:
			subQuery = fmt.Sprintf("SELECT id FROM category WHERE parent_category_id = %d", req.TopCategory)
		case 3:
			subQuery = fmt.Sprintf("SELECT id FROM category WHERE id = %d", req.TopCategory)
		default:
			return nil, status.Errorf(codes.InvalidArgument, "Level值无效")
		}

		tx = tx.Where(fmt.Sprintf("category_id IN (%s)", subQuery))
	}

	var count int64
	tx.Count(&count)

	var goods []model.Goods
	result := tx.Preload("Category").Preload("Brands").Scopes(Paginate(int(req.Pages), int(req.PagePerNums))).Find(&goods)
	if result.Error != nil {
		return nil, result.Error
	}

	var respList []*proto.GoodsInfoResponse

	for _, good := range goods {
		respList = append(respList, ModelToGoods(good))
	}

	return &proto.GoodsListResponse{
		Total: int32(count),
		Data:  respList,
	}, nil
}

func (g GoodsServer) BatchGetGoods(ctx context.Context, req *proto.BatchGoodsIdInfo) (*proto.GoodsListResponse, error) {
	var goods []model.Goods
	var respList []*proto.GoodsInfoResponse

	result := global.DB.Where("Id IN (?)", req.Id).Find(&goods)
	for _, good := range goods {
		respList = append(respList, ModelToGoods(good))
	}
	return &proto.GoodsListResponse{
		Total: int32(result.RowsAffected),
		Data:  respList,
	}, nil
}

func (g GoodsServer) CreateGoods(ctx context.Context, req *proto.CreateGoodsInfo) (*proto.GoodsInfoResponse, error) {
	return nil, nil

}
func (g GoodsServer) DeleteGoods(ctx context.Context, req *proto.DeleteGoodsInfo) (*proto.Empty, error) {
	return nil, nil

}
func (g GoodsServer) UpdateGoods(ctx context.Context, req *proto.CreateGoodsInfo) (*proto.Empty, error) {
	return nil, nil

}
func (g GoodsServer) GetGoodsDetail(ctx context.Context, req *proto.GoodInfoRequest) (*proto.GoodsInfoResponse, error) {
	return nil, nil

}

func ModelToGoods(good model.Goods) *proto.GoodsInfoResponse {
	return &proto.GoodsInfoResponse{
		Id:              good.ID,
		CategoryId:      good.CategoryID,
		Name:            good.Name,
		GoodsSn:         good.GoodsSn,
		ClickNum:        good.ClickNum,
		SoldNum:         good.SoldNum,
		FavNum:          good.FavNum,
		MarketPrice:     good.MarkPrice,
		ShopPrice:       good.ShopPrice,
		GoodsBrief:      good.GoodsBrief,
		ShipFree:        good.ShipFree,
		Images:          good.Images,
		DescImages:      good.DescImages,
		GoodsFrontImage: good.GoodsFrontImage,
		IsNew:           good.IsNew,
		IsHot:           good.IsHot,
		OnSale:          good.OnSale,
		Category: &proto.CategoryBriefInfoResponse{
			Id:   good.Category.ID,
			Name: good.Category.Name,
		},
		Brand: &proto.BrandInfoResponse{
			Id:   good.Brands.ID,
			Name: good.Brands.Name,
			Logo: good.Brands.Logo,
		},
	}
}
