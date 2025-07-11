package protectedcontroller

import (
	"net/http"

	loging "github.com/Samarth1015/expense/Loging"
	request "github.com/Samarth1015/expense/dto/Request"
	jwttoken "github.com/Samarth1015/expense/dto/jwtToken"
	"github.com/Samarth1015/expense/model"
	"github.com/Samarth1015/expense/postgres"
	"github.com/gin-gonic/gin"
)

func UpdateExpense(c *gin.Context) {
	var body request.UpdateExpenseReq
	if err := c.ShouldBindJSON(&body); err != nil {
		loging.Logger.Error("Invalid request body: ", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	// Extract claims from context
	claimsVal, exists := c.Get("claims")
	if !exists {
		loging.Logger.Error("Claims not found in context")
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	claims, ok := claimsVal.(jwttoken.Claims)
	if !ok {
		loging.Logger.Error("Invalid claims format")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	// Fetch expense from DB
	var dbData model.Expense
	err := postgres.Db.Preload("User").Where("id = ?", body.ExpenseId).First(&dbData).Error
	if err != nil {
		loging.Logger.Error("Error finding expense: ", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": "Expense not found"})
		return
	}

	// Check ownership
	if claims.ID != dbData.UserID {
		loging.Logger.Warn("User is not the owner of the expense")
		c.JSON(http.StatusUnauthorized, gin.H{"error": "You are not authorized to update this expense"})
		return
	}

	// Update fields
	dbData.Amount = float64(body.Amount)
	dbData.Date = body.Date
	dbData.Description = body.Description

	// Save updated expense
	if err := postgres.Db.Save(&dbData).Error; err != nil {
		loging.Logger.Error("Error updating expense: ", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update expense"})
		return
	}

	c.JSON(http.StatusAccepted, gin.H{"message": "Expense updated successfully"})
}
