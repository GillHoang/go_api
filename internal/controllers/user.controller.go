package controllers

import (
	"github.com/gillhoang/go_api/internal/services"
	"github.com/gillhoang/go_api/pkg/responses"
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
	c.JSON(200, responses.ResponseData{
		Code: 200,
		Message: "Success",
		Data: uc.userService.GetUserByID(uid),
	})
}
