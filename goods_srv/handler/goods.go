package handler

import (
	proto "GopherMall/goods_srv/proto/.GoodsProto"
)

type GoodsServer struct {
	proto.UnimplementedGoodsServer
}

//func (s *GoodsServer) GoodsList(context.Context, *proto.GoodsFilterRequest) (*proto.GoodsListResponse, error) {
//}

//func (s *GoodsServer) BatchGetGoods(context.Context, *proto.BatchGoodsIdInfo) (*proto.GoodsListResponse, error) {
//
//}
//func (s *GoodsServer) CreateGoods(context.Context, *proto.CreateGoodsInfo) (*proto.GoodsInfoResponse, error) {
//
//}
//func (s *GoodsServer) DeleteGoods(context.Context, *proto.DeleteGoodsInfo) (*proto.Empty, error) {
//
//}
//func (s *GoodsServer) UpdateGoods(context.Context, *proto.CreateGoodsInfo) (*proto.Empty, error) {
//
//}
//func (s *GoodsServer) GetGoodsDetail(context.Context, *proto.GoodInfoRequest) (*proto.GoodsInfoResponse, error) {
//
//}
