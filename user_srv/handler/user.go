package handler

import (
	"GopherMall/user_srv/global"
	"GopherMall/user_srv/model"
	"context"
	"crypto/sha256"
	"errors"
	"fmt"
	"github.com/anaskhan96/go-password-encoder"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
	"strings"
	"time"

	proto "GopherMall/user_srv/proto/.UserProto"
)

type UserServer struct {
	proto.UnimplementedUserServer
}

func ModelUserToResponse(user model.User) proto.UserInfoResponse {
	var birthday uint64
	if user.Birthday != nil {
		birthday = uint64(user.Birthday.Unix())
	}
	return proto.UserInfoResponse{
		Id:       user.ID,
		Password: user.Password,
		Mobile:   user.Mobile,
		NickName: user.NickName,
		Gender:   user.Gender,
		Role:     int32(user.Role),
		BirthDay: birthday,
	}
}

func Paginate(pageNum, pageSize int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if pageNum == 0 {
			pageNum = 1
		}

		switch {
		case pageSize > 100:
			pageSize = 100
		case pageSize <= 0:
			pageSize = 10
		}

		offset := (pageNum - 1) * pageSize
		return db.Offset(offset).Limit(pageSize)
	}
}

func (s *UserServer) GetUserList(ctx context.Context, req *proto.PageInfo) (*proto.UserListResponse, error) {
	var users []model.User
	userListResponse := make([]*proto.UserInfoResponse, 0)

	result := global.DB.Find(&users)
	if result.Error != nil {
		return nil, result.Error
	}

	err := global.DB.Scopes(Paginate(int(req.PageNum), int(req.PageSize))).Find(&users).Error
	if err != nil {
		return nil, err
	}

	for _, user := range users {
		userinfo := ModelUserToResponse(user)
		userListResponse = append(userListResponse, &userinfo)
	}

	return &proto.UserListResponse{
		Total: int32(result.RowsAffected),
		Data:  userListResponse,
	}, nil
}

func (s *UserServer) GetUserByMobile(ctx context.Context, request *proto.MobileRequest) (*proto.UserInfoResponse, error) {
	user, err := global.FindByMobile(request.GetMobile())
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, status.Errorf(codes.NotFound, "No Such User, Mobile: %s", request.GetMobile())
	}
	if err != nil {
		return nil, err
	}
	userInfoResponse := ModelUserToResponse(user)
	return &userInfoResponse, nil
}

func (s *UserServer) GetUserById(ctx context.Context, request *proto.IdRequest) (*proto.UserInfoResponse, error) {
	user, err := global.FindById(request.GetId())
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, status.Errorf(codes.NotFound, "No Such User, Id: %d", request.GetId())
	}
	if err != nil {
		return nil, err
	}
	userInfoResponse := ModelUserToResponse(user)
	return &userInfoResponse, nil
}

func (s *UserServer) CreateUser(ctx context.Context, request *proto.CreateUserInfo) (*proto.UserInfoResponse, error) {
	_, err := global.FindByMobile(request.GetMobile())
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, status.Errorf(codes.AlreadyExists, "Mobile Has Been Used: %s", request.GetMobile())
	}

	salt, encodedPwd := password.Encode(request.GetPassword(), &password.Options{
		SaltLen:      16,
		Iterations:   100,
		KeyLen:       32,
		HashFunction: sha256.New,
	})

	user := model.User{
		Mobile:   request.Mobile,
		Password: fmt.Sprintf("$pbkdf2-sha512$%s$%s", salt, encodedPwd),
		NickName: request.GetNickName(),
	}

	result := global.DB.Create(&user)
	if result.Error != nil {
		return nil, status.Errorf(codes.Internal, "Create User Failed: %v", result.Error)
	}

	userInfoResp := ModelUserToResponse(user)
	return &userInfoResp, nil
}

func (s *UserServer) UpdateUser(ctx context.Context, request *proto.UpdateUserInfo) (*proto.Empty, error) {
	user, err := global.FindById(request.GetId())
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, status.Errorf(codes.NotFound, "No Such User, Id: %d", request.GetId())
	}
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Find User By Id in UpdateUser Failed: %v", err)
	}

	birthDay := time.Unix(int64(request.GetBirthDay()), 0)
	user.NickName = request.GetNickName()
	user.Birthday = &birthDay
	user.Gender = request.GetGender()

	result := global.DB.Save(&user)
	if result.Error != nil {
		return nil, status.Errorf(codes.Internal, "Update User Failed: %v", result.Error)
	}

	return &proto.Empty{}, nil
}

func (s *UserServer) CheckPasswordInfo(ctx context.Context, request *proto.PasswordCheck) (*proto.CheckResponse, error) {
	passwordInfo := strings.Split(request.GetEncryptedPassword(), "$")
	check := password.Verify(request.GetPassword(), passwordInfo[2], passwordInfo[3], &password.Options{
		SaltLen:      16,
		Iterations:   100,
		KeyLen:       32,
		HashFunction: sha256.New,
	})

	return &proto.CheckResponse{Success: check}, nil
}
