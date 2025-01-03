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

func (g GoodsServer) BrandList(ctx context.Context, req *proto.BrandFilterRequest) (*proto.BrandListResponse, error) {
	var brands []model.Brands
	var BrandList []*proto.BrandInfoResponse
	rows := global.DB.Find(&brands).RowsAffected

	result := global.DB.Scopes(Paginate(int(req.Pages), int(req.PagePerNums))).Find(&brands)
	if result.Error != nil {
		return nil, result.Error
	}

	for _, brand := range brands {
		BrandList = append(BrandList, &proto.BrandInfoResponse{
			Id:   brand.ID,
			Name: brand.Name,
			Logo: brand.Logo,
		})
	}

	return &proto.BrandListResponse{
		Total: int32(rows),
		Data:  BrandList,
	}, nil
}

func (g GoodsServer) CreateBrand(ctx context.Context, req *proto.BrandRequest) (*proto.BrandInfoResponse, error) {
	var count int64

	if global.DB.Model(&model.Brands{}).Where("Name = ?", req.Name).Count(&count); count != 0 {
		fmt.Println(count)
		return nil, status.Errorf(codes.AlreadyExists, "商品名已存在")
	}

	brand := &model.Brands{
		Name: req.Name,
		Logo: req.Logo,
	}
	if global.DB.Save(brand).Error != nil {
		return nil, global.DB.Error
	}

	return &proto.BrandInfoResponse{
		Id:   brand.ID,
		Name: brand.Name,
		Logo: brand.Logo,
	}, nil
}

func (g GoodsServer) DeleteBrand(ctx context.Context, req *proto.BrandRequest) (*proto.Empty, error) {
	result := global.DB.Delete(&model.Brands{}, req.Id)
	if result.RowsAffected == 0 {
		return nil, status.Errorf(codes.NotFound, "无效ID")
	}
	return &proto.Empty{}, nil
}

func (g GoodsServer) UpdateBrand(ctx context.Context, req *proto.BrandRequest) (*proto.Empty, error) {
	var count int64

	if global.DB.Model(&model.Brands{}).Where("Name = ?", req.Name).Count(&count); count != 0 {
		fmt.Println(count)
		return nil, status.Errorf(codes.AlreadyExists, "商品名已存在")
	}

	brands := model.Brands{}
	if req.Name != "" {
		brands.Name = req.Name
	}

	if req.Logo != "" {
		brands.Logo = req.Logo
	}

	brands.ID = req.Id

	if err := global.DB.Updates(&brands).Error; err != nil {
		return nil, err
	}

	return &proto.Empty{}, nil
}
