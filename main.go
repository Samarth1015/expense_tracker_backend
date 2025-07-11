package main

import (
	// "fmt"
	// "os"

	"fmt"
	"log"

	loging "github.com/Samarth1015/expense/Loging"
	"github.com/Samarth1015/expense/client"
	"github.com/Samarth1015/expense/config"
	"github.com/Samarth1015/expense/middleware"
	"github.com/Samarth1015/expense/postgres"
	"github.com/Samarth1015/expense/routes/authenticate"
	"github.com/Samarth1015/expense/routes/protected"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

// func CORSMiddleware() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		c.Writer.Header().Set("Access-Control-Allow-Origin", "*") // or your frontend domain
// 		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
// 		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
// 		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
// 		c.Writer.Header().Set("Access-Control-Expose-Headers", "expense_token")

// 		// Handle preflight
// 		if c.Request.Method == "OPTIONS" {
// 			c.AbortWithStatus(204)
// 			return
// 		}

// 		c.Next()
// 	}
// }

func main() {

	db := postgres.Db
	fmt.Print(db)
	err := godotenv.Load()
	if err != nil {
		log.Print("error in env loadin")
	}
	// client.InitOpenAI()
	client.GeminiInit()
	loging.InitialiseLogger()

	router := gin.Default()
	// router.Use(CORSMiddleware())
	router.Use(middleware.CORSMiddleware())
	r := router.Group("/api")

	r.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, map[string]any{"message": "pong"})
	})

	authenticate.RegisterAuthRoute(r)
	protected.PromptRoute(r)
	protected.ExpenseRoute(r)

	command := fmt.Sprintf("server running on port %s and in %s mode", config.Config().Port, config.Config().Project)
	fmt.Println("------->", postgres.Db)

	fmt.Println(command)
	router.Run(":" + config.Config().Port)

}
