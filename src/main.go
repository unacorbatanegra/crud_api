package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-playground/validator"
	"github.com/labstack/echo"
)

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name" validate:"required"`
}

var users = []User{}

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	if err := cv.validator.Struct(i); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return nil
}

func main() {
	e := echo.New()
	e.GET("/", home)

	e.Validator = &CustomValidator{validator: validator.New()}

	e.GET("/users", getUsers)
	e.POST("/users", createUser)
	e.DELETE("/users/:id", deleteUser)
	e.Logger.Fatal(e.Start(":1323"))
}

func home(c echo.Context) error {
	return c.JSON(http.StatusOK, "test")

}

func getUsers(c echo.Context) error {
	return c.JSON(http.StatusOK, users)
}

func deleteUser(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}
	for i := 0; i < len(users); i++ {
		if users[i].ID == id {
			users = append(users[:i], users[i+1:]...)
			return c.String(http.StatusAccepted, "succesful")
		}
	}
	return c.NoContent(http.StatusNoContent)

}

func createUser(c echo.Context) error {
	fmt.Println(c.RealIP())
	u := &User{
		ID: len(users) + 1,
	}
	if err := c.Bind(u); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := c.Validate(u); err != nil {
		return err
	}

	users = append(users, *u)
	return c.JSON(http.StatusCreated, u)

}
