package services

import (
	"github.com/unacorbatanegra/crud_api/src/users/entities"
	"github.com/unacorbatanegra/crud_api/src/users/repositories"
)

type UserService interface {
	GetAll() []entities.User
	GetById(id uint) (entities.User, error)
	FindByName(name string) (entities.User, error)
}

type userService struct {
	postgresRepository repositories.UserPostgresRepositoryInterface
}

func NewUserService(repository repositories.UserPostgresRepositoryInterface) UserService {
	return &userService{
		postgresRepository: repository}

}

func (c *userService) GetAll() []entities.User {
	return c.postgresRepository.GetAll()
}

func (c *userService) GetById(id uint) (entities.User, error) {
	return c.postgresRepository.GetById(id)
}

func (c *userService) FindByName(name string) (entities.User, error) {
	return c.postgresRepository.FindByName(name)
}
