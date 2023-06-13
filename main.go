package main

import (
	"github.com/eminoz/pg/api"
	"github.com/eminoz/pg/db"
	"github.com/eminoz/pg/model"
	"github.com/eminoz/pg/service"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func Init() {

	db.ConnectToDB()

	d := db.GetDB()
	d.AutoMigrate(&model.User{})
}
func main() {
	Init()
	f := fiber.New()
	// CORS middleware configuration
	config := cors.Config{
		AllowOrigins:     "*",
		AllowMethods:     "GET,POST,PUT,DELETE",
		AllowHeaders:     "Origin,Content-Type,Accept,Authorization",
		AllowCredentials: true,
	}

	// CORS middleware
	f.Use(cors.New(config))
	s := service.NewUserService()
	a := api.NewUserApi(s)

	f.Post("/create", a.CreateUser)
	f.Delete("/delete/:id", a.DeleteUser)
	f.Get("/get/:id", a.FindUserById)
	f.Get("/getall", a.GetAll)
	f.Listen(":8081")

}
