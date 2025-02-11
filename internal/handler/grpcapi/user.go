package grpcapi

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
	"log/slog"
	pb "preproj/internal/handler/grpcapi/gen/user"
	"preproj/internal/models"
	"preproj/internal/service"
	"time"
)

type UserService struct {
	pb.UnimplementedUserServiceServer
	Services *service.Service
}

func (u *UserService) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Second*5)
	defer cancel()

	user := models.User{
		Name:  req.GetName(),
		Email: req.GetEmail(),
	}

	if user.Name == "" || user.Email == "" {
		slog.Info("Invalid data", slog.String("name", user.Name), slog.String("email", user.Email))
		return nil, status.Errorf(codes.InvalidArgument, "name or email are empty")
	}

	id, err := u.Services.User.Create(ctx, &user)
	if err != nil {
		slog.Error("failed to create user", slog.Any("error", err))
		return nil, status.Errorf(codes.Internal, "failed create user: %v", err)
	}
	return &pb.CreateUserResponse{
		Id: id,
	}, nil
}

func (u *UserService) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.GetUserResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Second*5)
	defer cancel()
	userID := req.GetId()

	user, err := u.Services.User.GetByID(ctx, userID)
	if err != nil {
		slog.Error("failed to get user", slog.Any("error", err))
		return nil, status.Errorf(codes.Internal, "failed to get user: %v", err)
	}

	return &pb.GetUserResponse{
		User: &pb.User{
			Id:        user.ID,
			Name:      user.Name,
			Email:     user.Email,
			CreatedAt: timestamppb.New(user.CreatedAt),
			UpdatedAt: timestamppb.New(user.UpdatedAt),
		},
	}, nil
}
func (u *UserService) GetAllUsers(ctx context.Context, req *pb.GetUsersRequest) (*pb.GetUsersResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Second*5)
	defer cancel()

	users, err := u.Services.User.GetAll(ctx)
	if err != nil {
		slog.Error("failed to get all users", slog.Any("error", err))
		return nil, status.Errorf(codes.Internal, "failed to get users: %v", err)
	}

	var pbUsers []*pb.User

	for _, user := range users {
		pbUsers = append(pbUsers, &pb.User{
			Id:        user.ID,
			Name:      user.Name,
			Email:     user.Email,
			CreatedAt: timestamppb.New(user.CreatedAt),
			UpdatedAt: timestamppb.New(user.UpdatedAt),
		})
	}
	if len(pbUsers) == 0 {
		slog.Info("users array is empty")
		return &pb.GetUsersResponse{Users: pbUsers}, nil
	}
	return &pb.GetUsersResponse{Users: pbUsers}, nil
}
func (u *UserService) DeleteUser(ctx context.Context, req *pb.DeleteUserRequest) (*pb.DeleteUserResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Second*5)
	defer cancel()
	userID := req.GetId()

	err := u.Services.User.Delete(ctx, userID)
	if err != nil {
		slog.Error("failed to delete user", slog.Any("error", err))
		return nil, status.Errorf(codes.Internal, "failed to delete user: %v", err)
	}
	return &pb.DeleteUserResponse{}, nil
}
func (u *UserService) UpdateUser(ctx context.Context, req *pb.UpdateUserRequest) (*pb.UpdateUserResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Second*5)
	defer cancel()
	pbUser := req.GetUser()

	user := models.User{
		ID:    pbUser.Id,
		Name:  pbUser.Name,
		Email: pbUser.Email,
	}

	userID, err := u.Services.User.Update(ctx, user)
	if err != nil {
		slog.Error("failed to update user", slog.Any("error", err))
		return nil, status.Errorf(codes.Internal, "failed to update user: %v", err)
	}

	return &pb.UpdateUserResponse{
		Id: userID,
	}, nil
}
