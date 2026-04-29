package services

import "github.com/gillhoang/go_api/internal/repos"

type UserService struct {
	userRepo *repos.UserRepo
}

func NewUserService() *UserService {
	return &UserService{
		userRepo: repos.NewUserRepo(),
	}
}

func (us *UserService) GetUserByID(id string) string {
	return us.userRepo.GetUserByID(id)
}
