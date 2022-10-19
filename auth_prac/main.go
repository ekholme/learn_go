package main

//see
//https://seefnasrul.medium.com/create-your-first-go-rest-api-with-jwt-authentication-in-gin-framework-dbe5bda72817
//or
//https://gist.github.com/mrcrilly/7703d630f9d589636d20b630245b6415

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

//a user
type User struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

var users []*User

//a struct that we get register input from
type RegisterInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

//register a user from URL parameters
//see https://github.com/gin-gonic/gin#querystring-parameters
func Register(c *gin.Context) {
	u := c.Query("username")
	pw := c.Query("password")

	//var user *User

	user := &User{
		Username: u,
		Password: pw,
	}

	//this sort of works, but it will keep duplicates
	users = append(users, user)

	c.JSON(http.StatusOK, gin.H{"user_list": users})
}

func main() {
	//just creating a basic server right now
	r := gin.Default()

	public := r.Group("/api")

	public.GET("/register", Register)

	r.Run(":8080")
}
