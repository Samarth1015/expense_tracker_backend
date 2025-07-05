package middleware

import (
	"fmt"
	"log"
	"os"

	"net/http"
	"strings"

	"github.com/clerk/clerk-sdk-go/v2"
	"github.com/clerk/clerk-sdk-go/v2/jwt"
	"github.com/gin-gonic/gin"
)

func JwtVerifyClerk() gin.HandlerFunc {

	return func(c *gin.Context) {
		fmt.Println("inside middleware")

		autHeader := c.GetHeader("Authorization")

		if autHeader == "" || !strings.HasPrefix(autHeader, "Bearer ") {
			log.Print("token not found")
			c.AbortWithStatusJSON(http.StatusUnauthorized, map[string]any{
				"message": "token missing",
			})
			return
		}

		tokenStr := strings.TrimSpace(strings.TrimPrefix(autHeader, "Bearer "))

		clerk.SetKey(os.Getenv("CLERK_SECRET_KEY"))

		_, err := jwt.Verify(c, &jwt.VerifyParams{Token: tokenStr})
		if err != nil {
			fmt.Println(err)
			c.AbortWithStatusJSON(http.StatusForbidden, map[string]any{
				"message": "invlaid token",
				"error":   err,
			})
			return
		}

		c.Next()

	}
}
