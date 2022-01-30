package repositories

import (
	"github.com/unacorbatanegra/crud_api/infraestructures/db"
	"github.com/unacorbatanegra/crud_api/src/users/entities"
)

type UserPostgresRepositoryInterface interface {
	GetAll() []entities.User
	GetById(id uint) entities.User
	FindByName(name string) entities.User
}

type userPostgresRepository struct {
	DB db.PostgresDB
}

func NewPostgresRepository(DB db.PostgresDB) UserPostgresRepositoryInterface {
	return &userPostgresRepository{
		DB: DB,
	}
}

func (c *userPostgresRepository) GetAll() []entities.User {
	var users []entities.User
	c.DB.DB().Find(&users)
	return users
}

func (c *userPostgresRepository) GetById(id uint) entities.User {
	var user entities.User
	c.DB.DB().First(&user, id)
	return user
}

func (c *userPostgresRepository) FindByName(name string) entities.User {
	var user entities.User
	c.DB.DB().Where("name = ?", name).First(&user)
	return user
}
