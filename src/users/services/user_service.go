package services

import (
	"github.com/unacorbatanegra/crud_api/src/users/entities"
	"github.com/unacorbatanegra/crud_api/src/users/repositories"
)

type UserService interface {
	GetAll() []entities.User
	GetById(id uint) entities.User
	FindByName(name string) entities.User
}

type userService struct {
	userPostgresRepository repositories.UserPostgresRepositoryInterface
}

func NewUserService(repository repositories.UserPostgresRepositoryInterface) UserService {
	return &userService{
		userPostgresRepository: repository}

}

func (c *userService) GetAll() []entities.User {
	return c.userPostgresRepository.GetAll()
}

func (c *userService) GetById(id uint) entities.User {
	return c.userPostgresRepository.GetById(id)
}

func (c *userService) FindByName(name string) entities.User {
	return c.userPostgresRepository.FindByName(name)
}
