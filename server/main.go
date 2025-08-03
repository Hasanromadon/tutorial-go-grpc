package main

import (
	"context"
	"fmt"
	pb "go-grpc/pb/user" // Import package yang dihasilkan dari file proto
	"log"
	"net"

	"google.golang.org/grpc" // Import library gRPC
)

// userServer mengimplementasikan interface UserServiceServer yang di-generate oleh protoc
type userServer struct {
	pb.UnimplementedUserServiceServer // Embedding untuk forward compatibility
}

// GetUser adalah implementasi RPC GetUser
// Fungsi ini akan dijalankan saat client memanggil method GetUser
func (s *userServer) GetUser(ctx context.Context, req *pb.UserRequest) (*pb.UserResponse, error) {
	log.Println("🎯 Menerima request untuk ID:", req.GetId())

	// Simulasi response dari server
	return &pb.UserResponse{
		Name: "Hasan",
		Age:  25,
	}, nil
}

func main() {
	// Membuka koneksi TCP di port 50051
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("❌ Gagal listen: %v", err)
	}

	// Membuat instance server gRPC
	grpcServer := grpc.NewServer()

	// Mendaftarkan service UserService ke dalam server gRPC
	pb.RegisterUserServiceServer(grpcServer, &userServer{})

	// Menjalankan server gRPC
	fmt.Println("🚀 gRPC Server running at :50051")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("❌ Gagal serve: %v", err)
	}
}
