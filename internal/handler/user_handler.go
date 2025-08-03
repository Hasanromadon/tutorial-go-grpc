package handler

import (
	"context"
	"go-grpc/internal/service"
	pb "go-grpc/pb/user"
)

type UserHandler struct {
	pb.UnimplementedUserServiceServer
	UserService service.UserService
}

func NewUserHandler(svc service.UserService) *UserHandler {
	return &UserHandler{UserService: svc}
}

func (h *UserHandler) GetUser(ctx context.Context, req *pb.UserRequest) (*pb.UserResponse, error) {
	user := h.UserService.GetUserByID(req.GetId())
	return &pb.UserResponse{
		Name: user.Name,
		Age:  user.Age,
	}, nil
}
