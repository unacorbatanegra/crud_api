package handlers

import (
	"net/http"
	"strconv"

	"github.com/gofiber/fiber"
	"github.com/unacorbatanegra/crud_api/internal/web"
	"github.com/unacorbatanegra/crud_api/src/users/dto"
	"github.com/unacorbatanegra/crud_api/src/users/entities"
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
	var user entities.User
	pid := ctx.Params("id")
	id, err := strconv.ParseUint(pid, 10, 32)
	if err != nil {
		web.JsonResponse(ctx, http.StatusBadRequest, err.Error(), nil)
		return
	}
	user = c.UserService.GetById(uint(id))
	web.JsonResponse(ctx, http.StatusOK, "", user)
}

func (c *userHandler) FindByName(ctx *fiber.Ctx) {
	var user entities.User
	name := ctx.Params("name")

	user = c.UserService.FindByName(name)
	web.JsonResponse(ctx, http.StatusOK, "", user)
}
