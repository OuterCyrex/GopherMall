package handler

import (
	"GopherMall/userop_srv/global"
	"GopherMall/userop_srv/model"
	FavProto "GopherMall/userop_srv/proto/.FavProto"
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type FavServer struct {
	FavProto.UnimplementedFavServer
}

func (f FavServer) GetFavList(ctx context.Context, req *FavProto.UserFavRequest) (*FavProto.UserFavListResponse, error) {
	var rsp FavProto.UserFavListResponse
	var userFavs []model.UserFav
	var userFavList []*FavProto.UserFavResponse

	result := global.DB.Where(&model.UserFav{User: req.UserId, Goods: req.GoodsId}).Find(&userFavs)
	rsp.Total = int32(result.RowsAffected)

	for _, userFav := range userFavs {
		userFavList = append(userFavList, &FavProto.UserFavResponse{
			UserId:  userFav.User,
			GoodsId: userFav.Goods,
		})
	}

	rsp.Data = userFavList

	return &rsp, nil
}

func (f FavServer) AddUserFav(ctx context.Context, req *FavProto.UserFavRequest) (*FavProto.Empty, error) {
	var userFav model.UserFav

	userFav.User = req.UserId
	userFav.Goods = req.GoodsId

	global.DB.Save(&userFav)

	return &FavProto.Empty{}, nil
}

func (f FavServer) DeleteUserFav(ctx context.Context, req *FavProto.UserFavRequest) (*FavProto.Empty, error) {
	if result := global.DB.Unscoped().Where("goods=? and user=?", req.GoodsId, req.UserId).Delete(&model.UserFav{}); result.RowsAffected == 0 {
		return nil, status.Errorf(codes.NotFound, "收藏记录不存在")
	}
	return &FavProto.Empty{}, nil
}

func (f FavServer) GetUserFavDetail(ctx context.Context, req *FavProto.UserFavRequest) (*FavProto.Empty, error) {
	var userFav model.UserFav
	if result := global.DB.Where("goods=? and user=?", req.GoodsId, req.UserId).Find(&userFav); result.RowsAffected == 0 {
		return nil, status.Errorf(codes.NotFound, "收藏记录不存在")
	}
	return &FavProto.Empty{}, nil
}
