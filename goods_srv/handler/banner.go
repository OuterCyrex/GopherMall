package handler

import (
	"GopherMall/goods_srv/global"
	"GopherMall/goods_srv/model"
	proto "GopherMall/goods_srv/proto/.GoodsProto"
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (g GoodsServer) BannerList(ctx context.Context, req *proto.Empty) (*proto.BannerListResponse, error) {
	var banners []model.Banner
	var BannerList []*proto.BannerResponse
	rows := global.DB.Find(&banners).RowsAffected

	result := global.DB.Find(&banners)
	if result.Error != nil {
		return nil, result.Error
	}

	for _, brand := range banners {
		BannerList = append(BannerList, &proto.BannerResponse{
			Id:    brand.ID,
			Index: brand.Index,
			Image: brand.Image,
			Url:   brand.Url,
		})
	}

	return &proto.BannerListResponse{
		Total: int32(rows),
		Data:  BannerList,
	}, nil
}

func (g GoodsServer) CreateBanner(ctx context.Context, req *proto.BannerRequest) (*proto.BannerResponse, error) {
	banner := &model.Banner{
		Index: req.Index,
		Image: req.Image,
		Url:   req.Url,
	}
	if global.DB.Save(banner).Error != nil {
		return nil, global.DB.Error
	}

	return &proto.BannerResponse{
		Id:    banner.ID,
		Index: banner.Index,
		Image: banner.Image,
		Url:   banner.Url,
	}, nil
}

func (g GoodsServer) DeleteBanner(ctx context.Context, req *proto.BannerRequest) (*proto.Empty, error) {
	result := global.DB.Delete(&model.Banner{}, req.Id)
	if result.RowsAffected == 0 {
		return nil, status.Errorf(codes.NotFound, "无效ID")
	}
	return &proto.Empty{}, nil
}

func (g GoodsServer) UpdateBanner(ctx context.Context, req *proto.BannerRequest) (*proto.Empty, error) {
	banner := model.Banner{}
	if req.Index != 0 {
		banner.Index = req.Index
	}

	if req.Image != "" {
		banner.Image = req.Image
	}

	if req.Url != "" {
		banner.Url = req.Url
	}

	banner.ID = req.Id

	if err := global.DB.Updates(&banner).Error; err != nil {
		return nil, err
	}

	return &proto.Empty{}, nil
}
