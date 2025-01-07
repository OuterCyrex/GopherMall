package handler

import (
	"GopherMall/goods_srv/global"
	"GopherMall/goods_srv/model"
	proto "GopherMall/goods_srv/proto/.GoodsProto"
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (g GoodsServer) CategoryBrandList(ctx context.Context, req *proto.CategoryBrandFilterRequest) (*proto.CategoryBrandListResponse, error) {
	var count int64
	var categoryBrand []model.GoodsCategoryBrand
	global.DB.Model(model.GoodsCategoryBrand{}).Count(&count)

	err := global.DB.Preload("Brands").Preload("Category").Scopes(Paginate(int(req.Pages), int(req.PagePerNums))).Model(model.GoodsCategoryBrand{}).Find(&categoryBrand).Error
	if err != nil {
		return nil, err
	}

	var respList []*proto.CategoryBrandResponse

	for _, cb := range categoryBrand {
		respList = append(respList, &proto.CategoryBrandResponse{
			Id: cb.ID,
			Brand: &proto.BrandInfoResponse{
				Id:   cb.BrandsID,
				Name: cb.Brands.Name,
				Logo: cb.Brands.Logo,
			},
			Category: &proto.CategoryInfoResponse{
				Id:             cb.CategoryID,
				Name:           cb.Category.Name,
				ParentCategory: cb.Category.ParentCategoryID,
				Level:          cb.Category.Level,
				IsTab:          cb.Category.IsTab,
			},
		})
	}

	return &proto.CategoryBrandListResponse{
		Total: int32(count),
		Data:  respList,
	}, nil
}

func (g GoodsServer) GetCategoryBrandList(ctx context.Context, req *proto.CategoryInfoRequest) (*proto.BrandListResponse, error) {

	var count int64

	var categoryBrand []model.GoodsCategoryBrand
	var brands []*proto.BrandInfoResponse

	r := global.DB.Preload("Brands").Model(model.GoodsCategoryBrand{}).Where("category_id = ?", req.Id).Find(&categoryBrand).Count(&count)
	if r.Error != nil {
		return nil, r.Error
	}

	for _, cb := range categoryBrand {
		brands = append(brands, &proto.BrandInfoResponse{
			Id:   cb.BrandsID,
			Name: cb.Brands.Name,
			Logo: cb.Brands.Logo,
		})
	}

	return &proto.BrandListResponse{
		Total: int32(count),
		Data:  brands,
	}, nil
}

func (g GoodsServer) CreateCategoryBrand(ctx context.Context, req *proto.CategoryBrandRequest) (*proto.CategoryBrandResponse, error) {
	var count int64

	if global.DB.Model(&model.Category{}).Where("id = ?", req.CategoryId).Count(&count); count == 0 {
		return nil, status.Errorf(codes.NotFound, "标签不存在")
	}

	if global.DB.Model(&model.Brands{}).Where("id = ?", req.BrandId).Count(&count); count == 0 {
		return nil, status.Errorf(codes.NotFound, "品牌不存在")
	}

	categoryBrand := model.GoodsCategoryBrand{
		CategoryID: req.CategoryId,
		BrandsID:   req.BrandId,
	}

	err := global.DB.Save(&categoryBrand).Error
	if err != nil {
		return nil, err
	}

	return &proto.CategoryBrandResponse{
		Id: categoryBrand.ID,
	}, nil
}

func (g GoodsServer) DeleteCategoryBrand(ctx context.Context, req *proto.CategoryBrandRequest) (*proto.Empty, error) {
	var categoryBrand model.GoodsCategoryBrand

	categoryBrand.ID = req.Id

	if r := global.DB.Delete(&categoryBrand); r.RowsAffected == 0 {
		return nil, status.Errorf(codes.NotFound, "无效ID")
	}

	return &proto.Empty{}, nil
}

func (g GoodsServer) UpdateCategoryBrand(ctx context.Context, req *proto.CategoryBrandRequest) (*proto.Empty, error) {
	var count int64

	var categoryBrand model.GoodsCategoryBrand
	if global.DB.Model(&model.GoodsCategoryBrand{}).Where("id = ?", req.Id).Count(&count); count == 0 {
		return nil, status.Errorf(codes.NotFound, "Category_Brand不存在")
	}

	categoryBrand.ID = req.Id
	if req.CategoryId != 0 {
		if global.DB.Model(&model.Category{}).Where("id = ?", req.CategoryId).Count(&count); count == 0 {
			return nil, status.Errorf(codes.NotFound, "标签不存在")
		}
		categoryBrand.CategoryID = req.CategoryId
	}

	if req.BrandId != 0 {
		if global.DB.Model(&model.Brands{}).Where("id = ?", req.BrandId).Count(&count); count == 0 {
			return nil, status.Errorf(codes.NotFound, "品牌不存在")
		}
		categoryBrand.BrandsID = req.BrandId
	}

	err := global.DB.Updates(&categoryBrand).Error
	if err != nil {
		return nil, err
	}

	return &proto.Empty{}, nil
}
