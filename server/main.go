package main

import (
	"context"
	"go-grpc/internal/handler"
	"go-grpc/internal/service"
	pb "go-grpc/pb/user"
	"log"
	"net"
	"time"

	"go.uber.org/zap"
	"google.golang.org/grpc"
)

var logger *zap.Logger

func init() {
	var err error
	logger, err = zap.NewProduction() // atau zap.NewDevelopment() untuk log berwarna
	if err != nil {
		log.Fatalf("cannot init zap: %v", err)
	}
}

func unaryZapInterceptor(
	ctx context.Context,
	req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler,
) (interface{}, error) {
	start := time.Now()
	res, err := handler(ctx, req)

	logger.Info("Unary gRPC Request",
		zap.String("method", info.FullMethod),
		zap.Duration("duration", time.Since(start)),
		zap.Bool("error", err != nil),
		zap.Error(err),
	)

	return res, err
}

func streamZapInterceptor(
	srv interface{},
	ss grpc.ServerStream,
	info *grpc.StreamServerInfo,
	handler grpc.StreamHandler,
) error {
	start := time.Now()
	err := handler(srv, ss)

	logger.Info("Stream gRPC Request",
		zap.String("method", info.FullMethod),
		zap.Duration("duration", time.Since(start)),
		zap.Bool("error", err != nil),
		zap.Error(err),
	)

	return err
}

func main() {
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("❌ Gagal listen: %v", err)
	}

	grpcServer := grpc.NewServer(
		grpc.UnaryInterceptor(unaryZapInterceptor),
		grpc.StreamInterceptor(streamZapInterceptor),
	)

	userService := service.NewUserService()
	userHandler := handler.NewUserHandler(userService)

	pb.RegisterUserServiceServer(grpcServer, userHandler)

	log.Println("🚀 gRPC server running on :50051")
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("❌ Gagal serve: %v", err)
	}
}
