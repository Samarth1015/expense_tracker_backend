package authenticate

import (
	authcontroller "github.com/Samarth1015/expense/controller/authController"
	"github.com/Samarth1015/expense/middleware"
	"github.com/gin-gonic/gin"
)

func RegisterAuthRoute(router *gin.RouterGroup) {

	auth := router.Group("/auth")

	auth.Use(middleware.JwtVerifyClerk())

	auth.POST("/login", authcontroller.Login)

}
