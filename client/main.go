package main

import (
	"context"
	"fmt"
	user "go-grpc/pb/user"
	"log"
	"time"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("❌ Gagal konek: %v", err)
	}
	defer conn.Close()

	client := user.NewUserServiceClient(conn)

	// Unary call
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	res, err := client.GetUser(ctx, &user.UserRequest{Id: 1})
	if err != nil {
		log.Fatalf("❌ Gagal GetUser: %v", err)
	}
	fmt.Printf("👤 [Unary] Nama: %s | Umur: %d\n", res.GetName(), res.GetAge())

	// Server Streaming call
	fmt.Println("📡 [Streaming] Daftar Pengguna:")
	stream, err := client.ListUsers(context.Background(), &user.Empty{})
	if err != nil {
		log.Fatalf("❌ Gagal ListUsers: %v", err)
	}

	for {
		user, err := stream.Recv()
		if err != nil {
			break
		}
		fmt.Printf("➡️ %s (%d tahun)\n", user.GetName(), user.GetAge())
	}
}
