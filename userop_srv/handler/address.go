package handler

import (
	"GopherMall/userop_srv/global"
	"GopherMall/userop_srv/model"
	AddressProto "GopherMall/userop_srv/proto/.AddressProto"
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type AddressServer struct {
	AddressProto.UnimplementedAddressServer
}

func (a AddressServer) GetAddressList(ctx context.Context, req *AddressProto.AddressRequest) (*AddressProto.AddressListResponse, error) {
	var addresses []model.Address
	var rsp AddressProto.AddressListResponse
	var addressResponse []*AddressProto.AddressResponse

	if result := global.DB.Where(&model.Address{User: req.UserId}).Find(&addresses); result.RowsAffected != 0 {
		rsp.Total = int32(result.RowsAffected)
	}

	for _, address := range addresses {
		addressResponse = append(addressResponse, &AddressProto.AddressResponse{
			Id:           address.ID,
			UserId:       address.User,
			Province:     address.Province,
			City:         address.City,
			District:     address.District,
			Address:      address.Address,
			SignerName:   address.SignerName,
			SignerMobile: address.SignerMobile,
		})
	}
	rsp.Data = addressResponse

	return &rsp, nil
}

func (a AddressServer) CreateAddress(ctx context.Context, req *AddressProto.AddressRequest) (*AddressProto.AddressResponse, error) {
	var address model.Address

	address.User = req.UserId
	address.Province = req.Province
	address.City = req.City
	address.District = req.District
	address.Address = req.Address
	address.SignerName = req.SignerName
	address.SignerMobile = req.SignerMobile

	global.DB.Save(&address)

	return &AddressProto.AddressResponse{Id: address.ID}, nil
}

func (a AddressServer) DeleteAddress(ctx context.Context, req *AddressProto.AddressRequest) (*AddressProto.Empty, error) {
	if result := global.DB.Where("id=? and user=?", req.Id, req.UserId).Delete(&model.Address{}); result.RowsAffected == 0 {
		return nil, status.Errorf(codes.NotFound, "收货地址不存在")
	}
	return &AddressProto.Empty{}, nil
}

func (a AddressServer) UpdateAddress(ctx context.Context, req *AddressProto.AddressRequest) (*AddressProto.Empty, error) {
	var address model.Address

	if result := global.DB.Where("id=? and user=?", req.Id, req.UserId).Find(&address); result.RowsAffected == 0 {
		return nil, status.Errorf(codes.NotFound, "收货地址不存在")
	}

	if req.Province != "" {
		address.Province = req.Province
	}
	if req.City != "" {
		address.City = req.City
	}
	if req.District != "" {
		address.District = req.District
	}
	if req.Address != "" {
		address.Address = req.Address
	}
	if req.SignerName != "" {
		address.SignerName = req.SignerName
	}
	if req.SignerMobile != "" {
		address.SignerMobile = req.SignerMobile
	}

	global.DB.Save(&address)

	return &AddressProto.Empty{}, nil
}
