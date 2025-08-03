package main

import (
	"context"
	"fmt"
	pb "go-grpc/pb/user" // 🟢 Import package hasil generate .proto (berisi stub gRPC)
	"log"
	"time"

	"google.golang.org/grpc" // 🟢 Library utama gRPC di Go
)

func main() {
	// 🟢 Membuka koneksi ke server gRPC
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure()) // ⚠️ grpc.Dial digunakan untuk membuat koneksi client → server gRPC
	if err != nil {
		log.Fatalf("❌ Gagal konek: %v", err)
	}
	defer conn.Close() // 🟢 Menutup koneksi gRPC saat selesai

	// 🟢 Membuat stub client dari service UserService (dari .proto)
	client := pb.NewUserServiceClient(conn)

	// Membuat context dengan timeout
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	// 🟢 Memanggil method RPC GetUser dari server gRPC
	res, err := client.GetUser(ctx, &pb.UserRequest{Id: 1}) // 🟢 client memanggil remote procedure GetUser
	if err != nil {
		log.Fatalf("❌ Gagal panggil GetUser: %v", err)
	}

	// Menampilkan response
	fmt.Printf("👤 User: %s, Age: %d\n", res.GetName(), res.GetAge()) // 🟢 Response hasil dari RPC server
}
