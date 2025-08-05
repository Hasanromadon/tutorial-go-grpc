package service

import "fmt"

type User struct {
	Name string
	Age  int32
}

type UserService interface {
	GetUserByID(id int32) *User
	GetAllUsers() []*User
	UploadUsers(users []*User) int32 // << new
}

type userServiceImpl struct{}

func NewUserService() UserService {
	return &userServiceImpl{}
}

func (s *userServiceImpl) GetUserByID(id int32) *User {
	if id == 1 {
		return &User{Name: "Andi", Age: 25}
	}
	return &User{Name: "Tidak Dikenal", Age: 0}
}

func (s *userServiceImpl) GetAllUsers() []*User {
	return []*User{
		{Name: "Andi", Age: 25},
		{Name: "Budi", Age: 30},
		{Name: "Citra", Age: 28},
	}
}

func (s *userServiceImpl) UploadUsers(users []*User) int32 {
	fmt.Printf("📥 Uploading %d users\n", len(users))
	for _, u := range users {
		fmt.Printf("✅ %s (%d tahun)\n", u.Name, u.Age)
	}
	return int32(len(users)) // hitung berapa yang berhasil
}
