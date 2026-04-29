package controllers

import (
	"github.com/gillhoang/hoc_api/internal/services"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	userService *services.UserService
}

func NewUserController() *UserController {
	return &UserController{
		userService: services.NewUserService(),
	}
}

func (uc *UserController) GetUser(c *gin.Context) {
	uid := c.Param("id")
	c.JSON(200, gin.H{
		"id":   uid,
		"name": uc.userService.GetUserByID(uid),
	})
}
