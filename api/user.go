package api

import (
	"github.com/eminoz/pg/model"
	"github.com/eminoz/pg/service"
	"github.com/gofiber/fiber/v2"
)

type UserApi interface {
	CreateUser(c *fiber.Ctx) error
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
	c.BodyParser(user)
	result := u.s.SaveUser(*user)
	return c.JSON(result)
}
