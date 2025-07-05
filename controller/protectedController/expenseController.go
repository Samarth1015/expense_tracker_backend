package protectedcontroller

import (
	"fmt"
	"time"

	loging "github.com/Samarth1015/expense/Loging"
	request "github.com/Samarth1015/expense/dto/Request"
	jwttoken "github.com/Samarth1015/expense/dto/jwtToken"
	"github.com/Samarth1015/expense/model"

	"github.com/Samarth1015/expense/postgres"

	"github.com/gin-gonic/gin"
)

func AddExpense(c *gin.Context) {
	var claims jwttoken.Claims
	var reqBody []request.ExpenseReq
	c.ShouldBindJSON(&reqBody)
	fmt.Println("------>", reqBody)
	var data []*model.Expense

	value, _ := c.Get("claims")
	claims = value.(jwttoken.Claims)

	for _, j := range reqBody {
		parsedDate, err := time.Parse("2006-01-02", j.Date)
		if err != nil {
			c.JSON(400, gin.H{"error": "Invalid date format. Use YYYY-MM-DD"})
			return
		}
		data = append(data, &model.Expense{
			Amount:      float64(j.Amount),
			Description: j.Description,
			UserID:      claims.ID,
			Date:        parsedDate,
		})
	}

	res := postgres.Db.Create(&data)
	if res.Error != nil {
		loging.Logger.Error("Error in making expense", res.Error.Error())
		c.Status(500)

		return
	}

	c.IndentedJSON(200, "Added Successfully")

}
