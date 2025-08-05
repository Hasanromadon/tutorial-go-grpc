package handler

import (
	"context"
	"fmt"
	"go-grpc/internal/service"
	pb "go-grpc/pb/user"
	"io"
	"log"
	"time"
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

func (h *UserHandler) UploadUsers(stream pb.UserService_UploadUsersServer) error {
	var users []*service.User

	for {
		req, err := stream.Recv()
		if err == io.EOF {
			count := h.UserService.UploadUsers(users)
			return stream.SendAndClose(&pb.UploadUserResponse{
				SuccessCount: count,
				Message:      fmt.Sprintf("%d users uploaded successfully", count),
			})
		}
		if err != nil {
			return err
		}

		users = append(users, &service.User{
			Name: req.GetName(),
			Age:  req.GetAge(),
		})
	}
}

func (h *UserHandler) UserChat(stream pb.UserService_UserChatServer) error {
	log.Println("📡 Bidirectional stream started...")

	for {
		msg, err := stream.Recv()
		if err == io.EOF {
			log.Println("🚫 Stream selesai oleh client")
			return nil
		}
		if err != nil {
			return err
		}

		log.Printf("📩 Pesan diterima dari %s: %s\n", msg.From, msg.Message)

		reply := &pb.ChatMessage{
			From:      "Server",
			Message:   fmt.Sprintf("Halo %s, pesanmu '%s' diterima!", msg.From, msg.Message),
			Timestamp: time.Now().Unix(),
		}

		if err := stream.Send(reply); err != nil {
			return err
		}
	}
}
