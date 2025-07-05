package protectedcontroller

import (
	"context"
	"encoding/json"
	"strings"

	"fmt"
	"net/http"

	loging "github.com/Samarth1015/expense/Loging"
	"github.com/Samarth1015/expense/client"
	request "github.com/Samarth1015/expense/dto/Request"
	promptdto "github.com/Samarth1015/expense/dto/promptdto"
	// "github.com/Samarth1015/expense/utils"

	// promptdto "github.com/Samarth1015/expense/dto/promptDTO"
	"github.com/gin-gonic/gin"
	"google.golang.org/genai"
)

func PromptController(c *gin.Context) {
	var res request.PromptRequest
	if err := c.ShouldBindJSON(&res); err != nil {
		loging.Logger.Error("error binding JSON: ", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	fullPrompt := fmt.Sprintf(`
You are an assistant for an expense tracker.
Extract structured data in JSON format from the following sentence.

Sentence: "%s"

Return only this JSON format in array only:
{
	"amount": number,
	"category": string,
	"date": string in YYYY-MM-DD,
	"description": string
}`, res.Prompt)

	result, err := client.Geminiclient.Models.GenerateContent(
		context.Background(),
		"gemini-2.0-flash",
		genai.Text(fullPrompt),
		nil,
	)
	if err != nil {
		loging.Logger.Error("error calling Gemini API: ", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to process prompt"})
		return
	}
	var jsonRes []promptdto.ExpenseData
	responseText := result.Text()
	// fmt.Println("--------------this is markup", responseText)
	responseText = strings.TrimPrefix(responseText, "```json")
	responseText = strings.TrimSuffix(responseText, "```")

	json.Unmarshal([]byte(responseText), &jsonRes)
	// fmt.Println("this is json:", jsonRes)
	// fmt.Println("---->jsondto", jsonRes.Amount, jsonRes.Category, jsonRes.Date, jsonRes.Description)

	if responseText == "" {
		loging.Logger.Error("empty response from Gemini API")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Empty response from API"})
		return
	}

	// Return the extracted JSON to the client
	c.JSON(http.StatusOK, gin.H{
		"data": jsonRes,
	})
}
