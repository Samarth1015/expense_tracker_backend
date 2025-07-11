package protectedcontroller

import (
	"net/http"

	loging "github.com/Samarth1015/expense/Loging"
	jwttoken "github.com/Samarth1015/expense/dto/jwtToken"
	"github.com/Samarth1015/expense/model"
	"github.com/Samarth1015/expense/postgres"
	"github.com/gin-gonic/gin"
)

func DeleteExpense(c *gin.Context) {
	expenseID := c.Param("expense_id")
	var data *model.Expense
	var claims jwttoken.Claims
	value, _ := c.Get("claims")
	claims = value.(jwttoken.Claims)
	err := postgres.Db.Where("id=?", expenseID).First(&data)
	// loging.Logger.Warn(data)
	if err.Error != nil {
		loging.Logger.Error("Id not found or not correct ", err.Error.Error())
		c.Status(http.StatusBadRequest)
		c.Abort()
		return
	}

	if claims.ID != data.UserID {
		c.Status(http.StatusUnauthorized)
		loging.Logger.Error("you are not owner of this expense")
		c.Abort()
		return

	}
	postgres.Db.Delete(&data)
	c.Status(http.StatusAccepted)

}
