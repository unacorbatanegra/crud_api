package entities

type User struct {
	ID   int    `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string 
}

func (t *User) TableName() string {
	return "users"
}
