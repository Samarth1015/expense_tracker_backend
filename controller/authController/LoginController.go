package authcontroller

import (
	"log"
	"net/http"
	"strings"
	"time"

	loging "github.com/Samarth1015/expense/Loging"
	jwttoken "github.com/Samarth1015/expense/dto/jwtToken"
	"github.com/Samarth1015/expense/model"
	"github.com/Samarth1015/expense/postgres"
	"github.com/Samarth1015/expense/service"

	"github.com/gin-gonic/gin"
)

type loginReq struct {
	UserId   string `json:"userId"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Role     string `json:"role,omitempty"`
}

var body loginReq

func Login(c *gin.Context) {
	// fmt.Print("inside login ")
	c.BindJSON(&body)
	var user model.User
	if !strings.HasPrefix(body.UserId, "user_") {
		loging.Logger.Error("Invalid User_id")
		c.Status(http.StatusBadRequest)

		c.Abort()
		return

	}
	c.Header("Content-Type", "application/json")
	// fmt.Println("------>calling how many timeL", body.Email)
	result := postgres.Db.First(&user, "id=?", body.UserId)
	// log.Print("-------->", result.Error.Error())
	if result.Error != nil {

		newUser := model.User{ID: body.UserId, Email: body.Email, Role: body.Role, UserName: body.Username, CreatedAt: time.Now()}
		createResult := postgres.Db.Create(&newUser)
		if createResult.Error != nil {
			log.Println(createResult.Error.Error())
			c.IndentedJSON(http.StatusInternalServerError, "Error in creating user")
			return

		}
		token, err := service.CreateToken(jwttoken.JwtToken{Id: body.UserId, Email: body.Email, Username: body.Username, Role: body.Role})
		if err != nil {
			c.IndentedJSON(http.StatusInternalServerError, "error in creating token ")
			return
		}
		if err := service.SendMail(body.Email); err != nil {
			loging.Logger.Warn("err", err)
			c.Abort()
		}
		c.Header("Access-Control-Expose-Headers", "expense_token")
		c.Header("expense_token", token)

		c.Status(http.StatusCreated)
		return

	}

	token, err := service.CreateToken(jwttoken.JwtToken{Id: user.ID, Email: user.Email, Username: user.UserName, Role: user.Role})
	if err != nil {
		log.Print(err)
		c.Status(http.StatusInternalServerError)
		return

	}
	c.Header("Access-Control-Expose-Headers", "expense_token")
	c.Header("expense_token", token)
	c.Status(http.StatusAccepted)

}
