package api

import (
	"fmt"
	"strconv"

	"github.com/eminoz/pg/model"
	"github.com/eminoz/pg/service"
	"github.com/gofiber/fiber/v2"
)

type UserApi interface {
	CreateUser(c *fiber.Ctx) error
	DeleteUser(c *fiber.Ctx) error
	FindUserById(c *fiber.Ctx) error
	GetAll(c *fiber.Ctx) error
}
type userApi struct {
	s service.UserService
}

func NewUserApi(s service.UserService) UserApi {
	return &userApi{
		s: s,
	}
}

func (u userApi) CreateUser(c *fiber.Ctx) error {
	user := &model.User{}
	fmt.Print(user)
	c.BodyParser(user)
	result := u.s.SaveUser(*user)
	return c.JSON(result)
}

func (u userApi) DeleteUser(c *fiber.Ctx) error {
	id := c.Params("id")
	err := u.s.DeleteUser(id)
	return err
}
func (u userApi) FindUserById(c *fiber.Ctx) error {
	id := c.Params("id")
	number, _ := strconv.ParseUint(id, 10, 0)

	res := u.s.FindUserById(uint(number))
	return c.JSON(res)

}

func (u userApi) GetAll(c *fiber.Ctx) error {
	all := u.s.GetAll()
	return c.JSON(all)

}
