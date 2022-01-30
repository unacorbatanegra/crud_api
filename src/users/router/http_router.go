package router

import (
	"log"

	"github.com/unacorbatanegra/crud_api/src/users/handlers"
	"github.com/unacorbatanegra/crud_api/src/users/repositories"
	"github.com/unacorbatanegra/crud_api/src/users/services"
)

func NewHttpRoute(structs RouterStruct) RouterStruct {
	log.Println("Setup HTTP Users Route")

	return structs
}

func (r *RouterStruct) GetRoute() {

	repo := repositories.NewPostgresRepository(r.PostgresDB)

	service := services.NewUserService(repo)
	handler := handlers.NewHttpHandler(service)

	r.Web.Get("/users", handler.GetAll)
	/// TODO: enable 
	// r.Web.Get("/users/:name", handler.FindByName)
	r.Web.Get("/users/:id", handler.GetById)
}
