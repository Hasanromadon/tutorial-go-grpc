package main

import (
	"context"
	"fmt"
	"go-fundamentals/helper" // Package custom berisi fungsi bantuan
	"go-fundamentals/models" // Package custom berisi model data
	"log"
	"os"
	"time"
)

func main() {
	fmt.Println("== Go Fundamentals Demo ==")

	// A. Package & Function
	// Contoh penggunaan function dari package lain
	result := helper.Add(10, 5)
	fmt.Println("🧮 Hasil penjumlahan:", result)

	// B. Variable & Constant
	// Contoh deklarasi variabel dan konstanta
	variableDemo()

	// C. Pointer & Error
	// Menunjukkan bagaimana pointer bekerja dan bagaimana error ditangani
	pointerDemo()
	errorHandlingDemo()

	// D. Struct & Interface
	// Menunjukkan penggunaan struct dan method pada struct
	structDemo()

	// E. Goroutine & Channel
	// Demonstrasi concurrency menggunakan goroutine dan channel
	concurrencyDemo()

	// F. Context
	// Contoh penggunaan context untuk mengatur timeout/cancel
	contextDemo()

	// G. Logging & Environment Variable
	// Contoh logging dan membaca environment variable
	loggingAndEnvDemo()
}

func variableDemo() {
	const appName = "GoApp"   // konstanta tidak bisa diubah
	var name string = "Hasan" // deklarasi variabel dengan tipe eksplisit
	age := 25                 // deklarasi variabel dengan type inference
	fmt.Printf("👋 Halo %s, umurmu %d, app: %s\n", name, age, appName)
}

func pointerDemo() {
	var x int = 10
	var p *int = &x                     // pointer menunjuk ke alamat memori variabel x
	fmt.Println("🔧 Nilai pointer:", *p) // dereference pointer, ambil nilai x lewat pointer
}

func errorHandlingDemo() {
	var data *int // pointer belum diinisialisasi, nil
	if data == nil {
		fmt.Println("❗ Error: pointer masih nil") // contoh error handling sederhana
	}
}

func structDemo() {
	// Membuat instance dari struct User (lihat di folder models/user.go)
	user := models.User{Name: "Andi", Age: 30}
	fmt.Println("🧍‍♂️", user.Greet()) // Memanggil method Greet() dari struct
}

func concurrencyDemo() {
	ch := make(chan string) // membuat channel string

	// menjalankan fungsi dalam goroutine (fungsi paralel)
	go func() {
		time.Sleep(1 * time.Second)
		ch <- "⏱️ Selesai dijalankan async" // kirim pesan ke channel setelah delay
	}()

	// menunggu hasil dari goroutine lewat channel
	fmt.Println(<-ch)
}

func contextDemo() {
	// membuat context dengan timeout 2 detik
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel() // pastikan cancel dipanggil untuk release resource

	select {
	case <-ctx.Done(): // akan terpanggil ketika context timeout
		fmt.Println("⛔ Context selesai:", ctx.Err()) // tampilkan error (misal: deadline exceeded)
	}
}

func loggingAndEnvDemo() {
	os.Setenv("APP_ENV", "development")      // set environment variable
	env := os.Getenv("APP_ENV")              // baca environment variable
	log.Println("🚀 Environment aktif:", env) // cetak log menggunakan package log
}
