package users

import (
	"react-auth/backend/domain/users"
	"react-auth/backend/services"
	"react-auth/backend/utils/errors"

	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {
	var user users.User

	if err := c.ShouldBindJSON(&user); err != nil {
		err := errors.NewBadRequestError("invalid json body")
		c.JSON(err.Status, err)
		return
	}

	services.CreateUser(user)

}
