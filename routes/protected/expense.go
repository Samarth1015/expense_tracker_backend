package protected

import (
	protectedcontroller "github.com/Samarth1015/expense/controller/protectedController"
	"github.com/Samarth1015/expense/middleware"

	"github.com/gin-gonic/gin"
)

func ExpenseRoute(router *gin.RouterGroup) {

	r := router.Group("/protected/expense")
	r.Use(middleware.JwtVerify())

	// r.GET("/",protectedcontroller.)
	r.GET("/ping/", func(c *gin.Context) {
		c.IndentedJSON(200, map[string]string{"msg": "pinging"})
	})

	r.POST("", protectedcontroller.AddExpense)
	r.GET("", protectedcontroller.GetAllExpense)
	r.PUT("", protectedcontroller.UpdateExpense)
	// r.DELETE("/",protectedcontroller.DeleteExpense);
	r.DELETE("/:expense_id", protectedcontroller.DeleteExpense)

}
