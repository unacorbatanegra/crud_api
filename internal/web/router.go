package web

import (
	"github.com/gofiber/fiber"
	"github.com/unacorbatanegra/crud_api/infraestructures/db"
)

type RouterStruct struct {
	Web        *fiber.App
	PostgresDB db.PostgresDB
}
