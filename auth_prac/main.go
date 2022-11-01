package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

//see
//https://www.sohamkamani.com/golang/jwt-authentication/
//or
//https://seefnasrul.medium.com/create-your-first-go-rest-api-with-jwt-authentication-in-gin-framework-dbe5bda72817
//or
//https://gist.github.com/mrcrilly/7703d630f9d589636d20b630245b6415
//or
//https://blog.logrocket.com/jwt-authentication-go/

// a user
type User struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

var privateStuff = map[string]string{
	"erice":   "message 1",
	"kendall": "message 2",
}

var users []*User

//handlers
func indexHandler(c *gin.Context) {
	data := gin.H{
		"hello": "world",
	}

	c.JSON(http.StatusOK, data)
}

func Register(c *gin.Context) {
	u := User{}

	c.ShouldBindJSON(&u)
}

func Login(c *gin.Context) {
	//TODO
}

//main function
func main() {
	r := gin.Default()

	r.GET("/", indexHandler)
	r.GET("/register", Register)
	r.GET("/login", Login)

	r.Run(":8080")
}
