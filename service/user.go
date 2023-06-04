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
	var user model.User

	result2 := u.db.Where("name = ?", usr.Name).First(&user)
	if result2.Error != nil {
		if result2.Error == gorm.ErrRecordNotFound {
			fmt.Println("User not found!")
		} else {
			log.Fatal(result2.Error)
		}
	}

	// Process the query result
	fmt.Println("ID:", user.ID, "Name:", user.Name)
	return user

}
