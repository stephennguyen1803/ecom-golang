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
	GetUserByName(name string) string
}

func NewUserController(UserService UserServiceInterface) *UserControler {
	return &UserControler{UserService: UserService}
}

// controller -> service -> repo -> model -> dbs
func (uc *UserControler) GetUserName(c *gin.Context) {
	name := c.Param("name")
	c.JSON(http.StatusOK, gin.H{ // map string
		"message": "Helllo" + uc.UserService.GetUserByName(name),
	})
}
