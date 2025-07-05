package protectedcontroller

import (
	"net/http"

	jwttoken "github.com/Samarth1015/expense/dto/jwtToken"
	"github.com/Samarth1015/expense/model"
	"github.com/Samarth1015/expense/postgres"
	"github.com/gin-gonic/gin"
)

func GetAllExpense(c *gin.Context) {
	var claims jwttoken.Claims

	value, exist := c.Get("claims")
	if !exist {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}
	cClaims, ok := value.(jwttoken.Claims)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token claims"})
		return
	}
	claims = cClaims

	var allExpense []model.Expense
	result := postgres.Db.
		Where("user_id = ?", claims.ID).
		Preload("User").
		Find(&allExpense)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch expenses"})
		return
	}

	c.JSON(http.StatusOK, allExpense)
}
