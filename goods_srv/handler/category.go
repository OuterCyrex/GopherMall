package handler

import (
	"GopherMall/goods_srv/global"
	"GopherMall/goods_srv/model"
	proto "GopherMall/goods_srv/proto/.GoodsProto"
	"context"
	"encoding/json"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (g GoodsServer) GetAllCategorysList(ctx context.Context, req *proto.Empty) (*proto.CategoryListResponse, error) {
	var categories []model.Category
	result := global.DB.Preload("SubCategory.SubCategory").Where("level = ?", 1).Find(&categories)
	if result.Error != nil {
		return nil, result.Error
	}

	b, _ := json.Marshal(categories)
	return &proto.CategoryListResponse{
		Total:    int32(result.RowsAffected),
		Data:     nil,
		JsonData: string(b),
	}, nil
}

func (g GoodsServer) GetSubCategory(ctx context.Context, req *proto.CategoryListRequest) (*proto.SubCategoryListResponse, error) {
	var respList []*proto.CategoryInfoResponse
	var info model.Category
	if count := global.DB.Model(&model.Category{}).Where("Id = ?", req.Id).Find(&info); count.RowsAffected == 0 {
		return nil, status.Errorf(codes.NotFound, "分类不存在")
	}

	var preloads string
	switch req.GetLevel() {
	case 1:
		preloads = "SubCategory.SubCategory"
	default:
		preloads = "SubCategory"
	}

	var SubCategories []model.Category
	result := global.DB.Preload(preloads).Where("parent_category_id = ?", req.GetId()).Find(&SubCategories)
	for _, cat := range SubCategories {
		respList = append(respList, &proto.CategoryInfoResponse{
			Id:             cat.ID,
			Name:           cat.Name,
			ParentCategory: cat.ParentCategoryID,
			Level:          cat.Level,
			IsTab:          cat.IsTab,
		})
	}

	return &proto.SubCategoryListResponse{
		Total: int32(result.RowsAffected),
		Info: &proto.CategoryInfoResponse{
			Id:             info.ID,
			Name:           info.Name,
			ParentCategory: info.ParentCategoryID,
			Level:          info.Level,
			IsTab:          info.IsTab,
		},
		SubCategorys: respList,
	}, nil
}

func (g GoodsServer) CreateCategory(ctx context.Context, req *proto.CategoryInfoRequest) (*proto.CategoryInfoResponse, error) {
	category := model.Category{
		Name:             req.Name,
		ParentCategoryID: req.ParentCategory,
		Level:            req.Level,
		IsTab:            req.IsTab,
	}

	if req.Level != 1 {
		var count int64
		if global.DB.Model(&model.Category{}).Where("ID = ?", req.ParentCategory).Count(&count); count == 0 {
			return nil, status.Errorf(codes.NotFound, "父分类不存在")
		}
	}

	if err := global.DB.Save(&category).Error; err != nil {
		return nil, err
	}

	return &proto.CategoryInfoResponse{Id: category.ID}, nil
}

func (g GoodsServer) DeleteCategory(ctx context.Context, req *proto.DeleteCategoryRequest) (*proto.Empty, error) {
	return nil, nil
}
func (g GoodsServer) UpdateCategory(ctx context.Context, req *proto.CategoryInfoRequest) (*proto.Empty, error) {
	return nil, nil
}
