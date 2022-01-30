package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gofiber/fiber"
	"github.com/unacorbatanegra/crud_api/internal/web"
	"github.com/unacorbatanegra/crud_api/src/users/dto"

	"github.com/unacorbatanegra/crud_api/src/users/services"
)

type UserHandlers interface {
	GetAll(ctx *fiber.Ctx)
	GetById(ctx *fiber.Ctx)
	FindByName(ctx *fiber.Ctx)
}

type userHandler struct {
	UserService services.UserService
}

func NewHttpHandler(userService services.UserService) UserHandlers {
	return &userHandler{
		UserService: userService,
	}
}

func (c *userHandler) GetAll(ctx *fiber.Ctx) {

	users := c.UserService.GetAll()
	response := dto.ParseFromEntities(users)
	web.JsonResponse(ctx, http.StatusOK, "", response)
}

func (c *userHandler) GetById(ctx *fiber.Ctx) {

	pid := ctx.Params("id")
	id, err := strconv.ParseUint(pid, 10, 32)
	if err != nil {
		web.JsonResponse(ctx, http.StatusBadRequest, err.Error(), nil)
		return
	}
	user, err := c.UserService.GetById(uint(id))
	if err != nil {
		fmt.Println(err)
	}
	web.JsonResponse(ctx, http.StatusOK, "", user)
}

func (c *userHandler) FindByName(ctx *fiber.Ctx) {

	name := ctx.Params("name")

	user, err := c.UserService.FindByName(name)
	if err != nil {
		fmt.Println(err)
	}
	web.JsonResponse(ctx, http.StatusOK, "", user)
}
