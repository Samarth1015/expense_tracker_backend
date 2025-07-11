package protected

import (
	protectedcontroller "github.com/Samarth1015/expense/controller/protectedController"
	"github.com/Samarth1015/expense/middleware"
	"github.com/gin-gonic/gin"
)

func PromptRoute(router *gin.RouterGroup) {

	r := router.Group("/protected/prompt")
	r.Use(middleware.JwtVerify())

	// r.GET("/",protectedcontroller.)
	r.GET("/ping", func(c *gin.Context) {
		c.IndentedJSON(200, map[string]string{"msg": "pinging"})
	})

	r.POST("", protectedcontroller.PromptController)

}
