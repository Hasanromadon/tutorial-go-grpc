package main

import (
	"go-grpc/internal/handler"
	"go-grpc/internal/service"
	pb "go-grpc/pb/user"
	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("❌ Gagal listen: %v", err)
	}

	grpcServer := grpc.NewServer()

	userService := service.NewUserService()
	userHandler := handler.NewUserHandler(userService)

	pb.RegisterUserServiceServer(grpcServer, userHandler)

	log.Println("🚀 gRPC server running on :50051")
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("❌ Gagal serve: %v", err)
	}
}
