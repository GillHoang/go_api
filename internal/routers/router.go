package routers

import (
	c "github.com/gillhoang/hoc_api/internal/controllers"
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	r := gin.Default()

	v1 := r.Group("/api/v1")
	{
		v1.GET("/user/:id", c.NewUserController().GetUser)
	}

	return r
}
