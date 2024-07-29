package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserControler struct {
	UserService UserServiceInterface
}

// Define it for easy to test
type UserServiceInterface interface {
	GetUserSerivce() string
}

func NewUserController(UserService UserServiceInterface) *UserControler {
	return &UserControler{UserService: UserService}
}

// controller -> service -> repo -> model -> dbs
func (uc *UserControler) GetUserInfo(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{ // map string
		"message": "Helllo" + uc.UserService.GetUserSerivce(),
	})
}
