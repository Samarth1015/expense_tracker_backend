package middleware

import (
	"fmt"
	"net/http"
	"strings"

	loging "github.com/Samarth1015/expense/Loging"
	jwttoken "github.com/Samarth1015/expense/dto/jwtToken"
	"github.com/Samarth1015/expense/service"

	"github.com/gin-gonic/gin"

	"go.uber.org/zap"
)

func JwtVerify() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("------------------------------calling")
		authHeader := c.GetHeader("Authorization")

		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			loging.Logger.Error("jwt token not passed")
			c.Status(http.StatusUnauthorized)
			c.Abort()
			return

		}

		tokenStr := strings.TrimSpace(strings.TrimPrefix(authHeader, "Bearer "))

		err := service.VerifyJWTToken(tokenStr)
		if err != nil {
			loging.Logger.Error("invalid token or token expired")
			c.Status(http.StatusUnauthorized)
			c.Abort()
			return
		}
		var claims jwttoken.Claims
		claims, err = service.ClaimToken(tokenStr)
		if err != nil {
			loging.Logger.Error("error in claiming", zap.Error(err))
			c.Status(http.StatusInternalServerError)
			c.Abort()
			return

		}

		fmt.Println(claims)
		c.Set("claims", claims)
		c.Next()

	}
}
