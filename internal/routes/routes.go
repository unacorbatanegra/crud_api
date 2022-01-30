package routes

import (
	"log"

	"github.com/gofiber/fiber"

	"github.com/unacorbatanegra/crud_api/internal/web"
	userRoute "github.com/unacorbatanegra/crud_api/src/users/router"
	// userRoute "crud_api/src/users/router/router.go"
)

type RouterStruct struct {
	web.RouterStruct
}

func NewHttpRoute(r RouterStruct) RouterStruct {
	log.Println("Loading the HTTP Router")
	return r
}

func (c *RouterStruct) GetRoutes() {
	c.Web.Get("/", func(c *fiber.Ctx) {
		c.Send([]byte("Hello this is my first route in go fiber"))
	})

	webRouterConfig := web.RouterStruct{
		Web:        c.Web,
		PostgresDB: c.PostgresDB,
	}

	userRouterStruct := userRoute.RouterStruct{
		RouterStruct: webRouterConfig,
	}

	userRouter := userRoute.NewHttpRoute(userRouterStruct)
	userRouter.GetRoute()

	c.Web.Use(func(c *fiber.Ctx) {
		c.Status(fiber.StatusNotFound).SendString("Sorry")
	})
}
