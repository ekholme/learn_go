package main

//see
//https://seefnasrul.medium.com/create-your-first-go-rest-api-with-jwt-authentication-in-gin-framework-dbe5bda72817
//or
//https://gist.github.com/mrcrilly/7703d630f9d589636d20b630245b6415

import (
	"errors"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

// a user
type User struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

var users []*User

// a struct that we get register input from
type RegisterInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// register a user from URL parameters
// see https://github.com/gin-gonic/gin#querystring-parameters
func Register(c *gin.Context) {
	u := c.Query("username")
	pw := c.Query("password")

	user := &User{
		Username: u,
		Password: pw,
	}

	err := user.hashPw()

	//this isn't the best way to do this, but w/e for now
	if err != nil {
		log.Fatal("Failed to hash password")
	}

	//this sort of works, but it will keep duplicates
	//not really the point of this code to get rid of dups, though
	users = append(users, user)

	c.JSON(http.StatusOK, gin.H{"user_list": users})
}

// creating a function to hash the password
func (u *User) hashPw() error {
	hashedPass, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)

	if err != nil {
		return err
	}

	u.Password = string(hashedPass)

	return nil
}

func (u *User) checkUserExists() (int, error) {
	for i, v := range users {
		if u.Username == v.Username {
			return i, nil
		}
	}
	return 0, errors.New("user doesn't exist")
}

// allow login
func Login(c *gin.Context) {
	//get params from querystring
	u := c.Query("username")
	pw := c.Query("password")

	inpUser := &User{
		Username: u,
		Password: pw,
	}

	if len(users) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"message": "no users registered"})
		return
	}

	//retrieve index of user in slice of users
	ind, err := inpUser.checkUserExists()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	ku := users[ind]

	err = bcrypt.CompareHashAndPassword([]byte(ku.Password), []byte(pw))

	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": ku})
}

func main() {
	//just creating a basic server right now
	r := gin.Default()

	public := r.Group("/api")

	public.GET("/register", Register)
	public.GET("/login", Login)

	r.Run(":8080")
}

//next step is to build and save tokens
