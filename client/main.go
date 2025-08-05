package main

import (
	"context"
	"fmt"
	user "go-grpc/pb/user"
	"io"
	"log"
	"time"

	"google.golang.org/grpc"
)

// Server Streaming call
func uploadUsers(client user.UserServiceClient) {
	stream, err := client.UploadUsers(context.Background())
	if err != nil {
		log.Fatalf("❌ Gagal buka stream UploadUsers: %v", err)
	}

	users := []user.UserRequest{
		{Name: "Dewi", Age: 22},
		{Name: "Eko", Age: 27},
		{Name: "Fajar", Age: 31},
	}

	for _, u := range users {
		fmt.Printf("📤 Mengirim: %s (%d tahun)\n", u.Name, u.Age)
		if err := stream.Send(&u); err != nil {
			log.Fatalf("❌ Gagal kirim data: %v", err)
		}
	}

	resp, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("❌ Gagal menerima respon server: %v", err)
	}

	fmt.Printf("🎉 Server: %s (%d berhasil)\n", resp.GetMessage(), resp.GetSuccessCount())
}

func startChat(client user.UserServiceClient) {
	stream, err := client.UserChat(context.Background())
	if err != nil {
		log.Fatalf("❌ Gagal membuka chat: %v", err)
	}

	done := make(chan struct{})

	// Goroutine untuk menerima pesan dari server
	go func() {
		for {
			in, err := stream.Recv()
			if err == io.EOF {
				close(done)
				return
			}
			if err != nil {
				log.Fatalf("❌ Gagal terima pesan: %v", err)
			}

			fmt.Printf("💬 %s >> %s\n", in.From, in.Message)
		}
	}()

	// Kirim beberapa pesan ke server
	messages := []string{"Halo!", "Apa kabar?", "Test 123"}

	for _, text := range messages {
		err := stream.Send(&user.ChatMessage{
			From:      "Client",
			Message:   text,
			Timestamp: time.Now().Unix(),
		})
		if err != nil {
			log.Fatalf("❌ Gagal kirim pesan: %v", err)
		}
		time.Sleep(1 * time.Second)
	}

	_ = stream.CloseSend()
	<-done
}

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

	// Client Streaming call
	uploadUsers(client)

	// Bidirectional streaming
	startChat(client)
}
