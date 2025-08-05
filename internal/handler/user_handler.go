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

func (h *UserHandler) ListUsers(_ *pb.Empty, stream pb.UserService_ListUsersServer) error {
	users := h.UserService.GetAllUsers()

	for _, user := range users {
		err := stream.Send(&pb.UserResponse{
			Name: user.Name,
			Age:  user.Age,
		})
		if err != nil {
			return err
		}
	}

	return nil
}

func (h *UserHandler) GetUser(ctx context.Context, req *pb.UserRequest) (*pb.UserResponse, error) {
	user := h.UserService.GetUserByID(req.GetId())
	return &pb.UserResponse{
		Name: user.Name,
		Age:  user.Age,
	}, nil
}
