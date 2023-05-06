package service

import (
	"context"
	"strconv"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"

	"github.com/sjxiang/ddshop/user_srv/cmd/user/dal/db"
	"github.com/sjxiang/ddshop/user_srv/cmd/user/pack"
	"github.com/sjxiang/ddshop/user_srv/model"
	"github.com/sjxiang/ddshop/user_srv/pb"
)

type UserServiceImpl struct {
	pb.UnimplementedUserServer
}

func (impl *UserServiceImpl) GetUserList(ctx context.Context, req *pb.GetUserListRequest) (*pb.GetUserListReply, error) {

	var total int64
	var res []*model.User

	offset := GetPageOffset(req.PageInfo.PageNum, req.PageInfo.PageSize)
	limit := int(req.PageInfo.PageSize)

	res, total, err := db.MGetUser(ctx, limit, offset)
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	return &pb.GetUserListReply{
		Data: pack.UserInfoList(res),
		Total: total,
	}, nil

}

// 通过手机号码查询用户
func (impl *UserServiceImpl) GetUserByEmail(ctx context.Context, req *pb.GetUserByEmailRequest) (*pb.UserInfo, error) {
	var res *model.User	

	res, err := db.QueryUserByEmail(ctx, req.Email)
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}
	
	return pack.UserInfo(res), nil
}

// 通过用户 id 查询用户
func (impl *UserServiceImpl) GetUserById(ctx context.Context, req *pb.GetUserByIdRequest) (*pb.UserInfo, error) {
	var res []*model.User	
	res, err := db.QueryUserByID(ctx, req.Id)
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	return pack.UserInfoList(res)[0], nil
}


func (impl *UserServiceImpl) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.UserInfo, error) {

	var user model.User

	// 查询是否注册
	res := db.QueryUserByEmailPlus(ctx, req.Email, &user)
	if res.RowsAffected == 1 {
		return nil, status.Errorf(codes.AlreadyExists, "用户已经存在")
	}
	
	// 新建用户
	user.Mobile = req.Mobile
	user.Email = req.Email
	user.NickName = req.Nickname
	user.Password = hashPassowrd(strconv.Itoa(int(req.Password)))

	if err := user.Check(); err != nil {
		return nil, status.Errorf(codes.AlreadyExists, "注册失败" + err.Error())
	}

	if err := db.CreateUser(ctx, &user); err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	return pack.UserInfo(&user), nil
}

// 个人中心 - 更新用户
func (impl *UserServiceImpl) UpdateUser(ctx context.Context, req *pb.UpdateUserRequest) (*emptypb.Empty, error) {

	
	if err := db.UpdateUser(ctx, req.Id, req.Mobile, req.Email, req.Nickname, req.Password, req.Birthday); err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	return&emptypb.Empty{}, nil
}


// 校验密码
func (impl *UserServiceImpl) VerifyPassword(ctx context.Context, req *pb.VerifyPasswordRequest) (*pb.VerifyPasswordReply, error) {
	return &pb.VerifyPasswordReply{
		Success: verifyPassowrd(req.PasswordHash, req.PlainText),
	}, nil
}

