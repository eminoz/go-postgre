package service

import (
	"fmt"
	"log"

	"github.com/eminoz/pg/db"
	"github.com/eminoz/pg/model"
	"gorm.io/gorm"
)

type UserService interface {
	SaveUser(usr model.User) model.User
	FindUserById(id uint) model.User
	DeleteUser(id string) error
	GetAll() []model.User
}
type userService struct {
	db *gorm.DB
}

func NewUserService() UserService {
	return &userService{
		db: db.GetDB(),
	}
}
func (u *userService) SaveUser(usr model.User) model.User {
	result := u.db.Create(&usr)

	if result.Error != nil {
		fmt.Println(result)
	}
	user := u.FindUserById(usr.ID)
	// Process the query result
	fmt.Println("ID:", user.ID, "Name:", user.Title)
	return user

}
func (u *userService) FindUserById(id uint) model.User {
	var user model.User

	result := u.db.Where("ID = ?", id).First(&user)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			fmt.Println("User not found!")
		} else {
			log.Fatal(result.Error)
		}
	}
	return user

}
func (u *userService) DeleteUser(id string) error {
	user := model.User{}
	result := u.db.Delete(&user, id)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (u *userService) GetAll() []model.User {
	var user []model.User
	u.db.Find(&user)
	return user
}
