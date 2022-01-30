package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber"
	"github.com/unacorbatanegra/crud_api/config"
	"github.com/unacorbatanegra/crud_api/infraestructures/db"
	"github.com/unacorbatanegra/crud_api/internal/routes"
	"github.com/unacorbatanegra/crud_api/internal/web"
)

func main() {

	port := config.Get().AppPort

	log.Println("Server runnig on PORT: ", port)

	app := fiber.New()

	postgresDb := db.NewPostgresClient()

	routeStruct := routes.RouterStruct{
		RouterStruct: web.RouterStruct{
			Web:        app,
			PostgresDB: postgresDb,
		},
	}
	router := routes.NewHttpRoute(routeStruct)
	router.GetRoutes()
	app.Listen(fmt.Sprintf(":%s", port))

}
