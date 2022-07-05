package userController

import (
	"jackthomson/go-api/models"
	userService "jackthomson/go-api/services/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetUser(c *gin.Context) {
	user, err := userService.GetUser(c.Param("name"))

	if err != nil {
		c.JSON(http.StatusBadRequest, models.Error{Message: err.Error()})
	} else {
		c.JSON(http.StatusOK, user)
	}
}

func GetUsers(c *gin.Context) {
	users, err := userService.GetUsers()

	if err != nil {
			c.JSON(http.StatusBadRequest, models.Error{Message: err.Error()})
	}

	c.JSON(http.StatusOK, users)
}

func CreateUser(c *gin.Context) {
	user, err := c.GetRawData()
	if err != nil {
		c.JSON(http.StatusBadRequest, models.Error{Message: err.Error()})
	}

	err = userService.CreateUser(user)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.Error{Message: err.Error()})
	}

	c.Status(http.StatusOK)
}