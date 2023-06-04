package main

import (
	"github.com/eminoz/pg/api"
	"github.com/eminoz/pg/db"
	"github.com/eminoz/pg/model"
	"github.com/eminoz/pg/service"
	"github.com/gofiber/fiber/v2"
)

func Init() {

	db.ConnectToDB()

	d := db.GetDB()
	d.AutoMigrate(&model.User{})
}
func main() {
	Init()
	f := fiber.New()
	s := service.NewUserService()
	a := api.NewUserApi(s)

	f.Post("/create", a.CreateUser)
	f.Listen(":8081")

}
