package service

type User struct {
	Name string
	Age  int32
}

type UserService interface {
	GetUserByID(id int32) *User
	GetAllUsers() []*User // << new
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
