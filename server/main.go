package main

import (
	"context"
	"go-grpc/internal/handler"
	"go-grpc/internal/service"
	pb "go-grpc/pb/user"
	"log"
	"net"
	"time"

	"google.golang.org/grpc"
)

func unaryLoggerInterceptor(
	ctx context.Context,
	req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler,
) (interface{}, error) {
	start := time.Now()

	// Proses request ke handler RPC
	res, err := handler(ctx, req)

	duration := time.Since(start)

	log.Printf("📝 Unary: %s | duration: %v | error: %v",
		info.FullMethod, duration, err)

	return res, err
}

func streamLoggerInterceptor(
	srv interface{},
	ss grpc.ServerStream,
	info *grpc.StreamServerInfo,
	handler grpc.StreamHandler,
) error {
	start := time.Now()

	// Proses stream ke handler
	err := handler(srv, ss)

	duration := time.Since(start)

	log.Printf("📡 Stream: %s | duration: %v | error: %v",
		info.FullMethod, duration, err)

	return err
}

func main() {
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("❌ Gagal listen: %v", err)
	}

	grpcServer := grpc.NewServer(
		grpc.UnaryInterceptor(unaryLoggerInterceptor),
		grpc.StreamInterceptor(streamLoggerInterceptor),
	)

	userService := service.NewUserService()
	userHandler := handler.NewUserHandler(userService)

	pb.RegisterUserServiceServer(grpcServer, userHandler)

	log.Println("🚀 gRPC server running on :50051")
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("❌ Gagal serve: %v", err)
	}
}
