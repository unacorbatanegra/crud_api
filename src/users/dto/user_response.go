package dto

import "github.com/unacorbatanegra/crud_api/src/users/entities"

type UserResponseBody struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func ParseFromEntities(data []entities.User) []UserResponseBody {
	var users []UserResponseBody

	for _, d := range data {
		user := UserResponseBody{
			ID:    d.UserID,
			Name:  d.Name,
			Email: d.Email,
		}
		users = append(users, user)
	}
	return users
}
