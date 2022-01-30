package repositories

import (
	"github.com/unacorbatanegra/crud_api/infraestructures/db"
	"github.com/unacorbatanegra/crud_api/src/users/entities"
)

type UserPostgresRepositoryInterface interface {
	GetAll() []entities.User
	GetById(id uint) (entities.User, error)
	FindByName(name string) (entities.User, error)
}

type userPostgresRepository struct {
	db db.PostgresDB
}

func NewPostgresRepository(DB db.PostgresDB) UserPostgresRepositoryInterface {
	return &userPostgresRepository{
		db: DB,
	}
}

func (c *userPostgresRepository) GetAll() []entities.User {
	var users []entities.User
	c.db.DB().Find(&users)
	return users
}

func (c *userPostgresRepository) GetById(id uint) (entities.User, error) {
	var user entities.User
	err := c.db.DB().First(&user, id).Error

	return user, err
}

func (c *userPostgresRepository) FindByName(name string) (entities.User, error) {
	var user entities.User
	err := c.db.DB().Where("name = ?", name).First(&user).Error
	return user, err
}
