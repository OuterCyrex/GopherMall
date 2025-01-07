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
	var brand model.Brands
	var category model.Category

	result := global.DB.Model(&model.Brands{}).Where("Id = ?", req.BrandId).Find(&brand)
	if result.RowsAffected == 0 {
		return nil, status.Errorf(codes.NotFound, "无效品牌ID")
	}

	result = global.DB.Model(&model.Category{}).Where("Id = ?", req.CategoryId).Find(&category)
	if result.RowsAffected == 0 {
		return nil, status.Errorf(codes.NotFound, "无效标签ID")
	}

	good := model.Goods{
		Category:        category,
		CategoryID:      req.CategoryId,
		Brands:          brand,
		BrandsID:        req.BrandId,
		OnSale:          req.OnSale,
		ShipFree:        req.ShipFree,
		IsNew:           req.IsNew,
		IsHot:           req.IsHot,
		Name:            req.Name,
		GoodsSn:         req.GoodsSn,
		MarkPrice:       req.MarketPrice,
		ShopPrice:       req.ShopPrice,
		GoodsBrief:      req.GoodsBrief,
		Images:          req.Images,
		DescImages:      req.DescImages,
		GoodsFrontImage: req.GoodsFrontImage,
	}

	if err := global.DB.Save(&good).Error; err != nil {
		return nil, err
	}

	return &proto.GoodsInfoResponse{
		Id: good.ID,
	}, nil
}

func (g GoodsServer) DeleteGoods(ctx context.Context, req *proto.DeleteGoodsInfo) (*proto.Empty, error) {
	result := global.DB.Where("Id = ?", req.Id).Delete(&model.Goods{})

	if result.RowsAffected == 0 {
		return nil, status.Errorf(codes.NotFound, "无效ID")
	}

	return &proto.Empty{}, nil
}

func (g GoodsServer) UpdateGoods(ctx context.Context, req *proto.CreateGoodsInfo) (*proto.Empty, error) {
	var brand model.Brands
	var category model.Category
	var former model.Goods

	if result := global.DB.Model(&model.Goods{}).Where("Id = ?", req.Id).Find(&former); result.RowsAffected == 0 {
		return nil, status.Errorf(codes.NotFound, "无效商品ID")
	}

	goods := model.Goods{
		OnSale:          req.OnSale,
		ShipFree:        req.ShipFree,
		IsNew:           req.IsNew,
		IsHot:           req.IsHot,
		Name:            req.Name,
		GoodsSn:         req.GoodsSn,
		MarkPrice:       req.MarketPrice,
		ShopPrice:       req.ShopPrice,
		GoodsBrief:      req.GoodsBrief,
		Images:          req.Images,
		DescImages:      req.DescImages,
		GoodsFrontImage: req.GoodsFrontImage,
	}

	goods.ID = req.Id

	if req.BrandId > 0 {
		result := global.DB.Model(&model.Brands{}).Where("Id = ?", req.BrandId).Find(&brand)
		if result.RowsAffected == 0 {
			return nil, status.Errorf(codes.NotFound, "无效品牌ID")
		}
		goods.Brands = brand
		goods.BrandsID = req.BrandId
	}

	if req.CategoryId > 0 {
		result := global.DB.Model(&model.Category{}).Where("Id = ?", req.CategoryId).Find(&category)
		if result.RowsAffected == 0 {
			return nil, status.Errorf(codes.NotFound, "无效标签ID")
		}

		goods.Category = category
		goods.CategoryID = req.CategoryId
	}

	goodsMap := map[string]interface{}{}

	if req.ShipFree != former.ShipFree {
		goodsMap["ship_free"] = req.ShipFree
	}
	if req.IsNew != former.IsNew {
		goodsMap["is_new"] = req.IsNew
	}
	if req.IsHot != former.IsHot {
		goodsMap["is_hot"] = req.IsHot
	}
	if req.OnSale != former.OnSale {
		goodsMap["on_sale"] = req.OnSale
	}

	result := global.DB.Updates(&goods).Updates(goodsMap)
	if result.Error != nil {
		return nil, result.Error
	}

	return &proto.Empty{}, nil
}
func (g GoodsServer) GetGoodsDetail(ctx context.Context, req *proto.GoodInfoRequest) (*proto.GoodsInfoResponse, error) {
	var Good model.Goods

	result := global.DB.Where("Id = ?", req.Id).Find(&Good)
	if result.RowsAffected == 0 {
		return nil, status.Errorf(codes.NotFound, "无效ID")
	}

	return ModelToGoods(Good), nil
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
