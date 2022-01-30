package entities

type User struct {
	UserID   int
	Name     string
	Email    string
	Password string
}

func (t *User) TablName() string {
	return "users"
}
