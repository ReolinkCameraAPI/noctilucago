package controllers

import (
	"github.com/ReolinkCameraAPI/noctilucago/internal/pkg/database/models"
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func (ac *ApiController) Login(c *gin.Context) (interface{}, error) {

	var user models.User

	if err := c.ShouldBindJSON(&user); err != nil {
		return nil, jwt.ErrMissingLoginValues
	}

	dbUser, err := ac.db.UserReadByUsername(user.Username)

	if err != nil {
		return nil, jwt.ErrFailedAuthentication
	}

	err = bcrypt.CompareHashAndPassword([]byte(dbUser.Password), []byte(user.Password))

	if err != nil {
		return nil, jwt.ErrFailedAuthentication
	}

	return map[string]interface{}{
		"UUID":     dbUser.UUID,
		"Username": dbUser.Username,
	}, nil
}

func (ac *ApiController) UserCreate(c *gin.Context) {
	var user models.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(500, &GenericResponse{
			Status:  "error",
			Message: "Incorrect data payload sent",
		})
		return
	}

	dbUser, err := ac.db.UserCreate(user.Username, user.Password)

	if err != nil {
		c.JSON(500, &GenericResponse{
			Status:  "error",
			Message: "User could not be created",
		})
		return
	}

	c.JSON(200, map[string]interface{}{
		"UUID":     dbUser.UUID,
		"Username": dbUser.Username,
	})
}

func (ac *ApiController) UserUpdate(c *gin.Context) {
	var user models.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(500, &GenericResponse{
			Status:  "error",
			Message: "Incorrect data payload sent",
		})
		return
	}

	dbUser, err := ac.db.UserUpdate(user.UUID, user.Username, user.Password)

	if err != nil {
		c.JSON(500, &GenericResponse{
			Status:  "error",
			Message: "Could not update user",
		})
		return
	}

	c.JSON(200, map[string]interface{}{
		"UUID":     dbUser.UUID,
		"Username": dbUser.Username,
	})
}

func (ac *ApiController) UserDelete(c *gin.Context) {
	var user models.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(500, &GenericResponse{
			Status:  "error",
			Message: "Incorrect data payload sent",
		})
		return
	}

	_, err := ac.db.UserDelete(user.UUID)

	if err != nil {
		c.JSON(500, &GenericResponse{
			Status:  "error",
			Message: "Could not delete user",
		})
		return
	}

	c.JSON(200, &GenericResponse{
		Status:  "success",
		Message: "User deleted",
	})
}

func (ac *ApiController) UserRead(c *gin.Context) {
	users, err := ac.db.UserRead()

	if err != nil {
		c.JSON(500, &GenericResponse{
			Status:  "error",
			Message: "Could not find users",
		})
		return
	}

	c.JSON(200, users)
}
