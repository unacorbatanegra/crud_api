package dto

import "github.com/unacorbatanegra/crud_api/src/users/entities"

type UserResponseBody struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func ParseFromEntity(data entities.User) UserResponseBody {
	return UserResponseBody{
		ID:    data.ID,
		Name:  data.Name,
		Email: data.Email,
	}
}

func ParseFromEntities(data []entities.User) []UserResponseBody {
	var users []UserResponseBody

	for _, d := range data {
		user := UserResponseBody{
			ID:    d.ID,
			Name:  d.Name,
			Email: d.Email,
		}
		users = append(users, user)
	}
	return users
}
