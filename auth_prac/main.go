package main

import (
	"errors"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
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

//custom claims
type CustomClaims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

var privateStuff = map[string]string{
	"erice":   "message 1",
	"kendall": "message 2",
}

var users []*User

// handlers
func indexHandler(c *gin.Context) {
	data := gin.H{
		"hello": "world",
	}

	c.JSON(http.StatusOK, data)
}

func Register(c *gin.Context) {
	u := User{}

	err := c.ShouldBindJSON(&u)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	//add password hash
	err = u.hashPass()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "failure hashing password"})
		return
	}

	//this won't handle duplicate users, but that's fine for now
	users = append(users, &u)

	c.JSON(http.StatusOK, gin.H{
		"message": "successfully registered",
		"user":    u,
	})

}

func Login(c *gin.Context) {
	//TODO

	u := User{}

	err := c.ShouldBindJSON(&u)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	//loop over user slice to ensure user exists
	ind, err := u.checkUserExists(users)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	ref := users[ind]

	//check that passwords match
	err = u.checkPwMatch(ref)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	tokenStr, err := u.generateToken()

	//RESUME HERE

}

//password hashing func
func (u *User) hashPass() error {

	hp, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)

	if err != nil {
		return err
	}

	u.Password = string(hp)

	return nil
}

//check that a user exists
func (u *User) checkUserExists(us []*User) (int, error) {
	for i, v := range us {
		if u.Username == v.Username {
			return i, nil
		}
	}
	return 0, errors.New("user doesn't exist")
}

//check that passwords match
func (u *User) checkPwMatch(r *User) error {

	rp := []byte(r.Password)
	up := []byte(u.Password)
	err := bcrypt.CompareHashAndPassword(rp, up)

	return err
}

//generate a token for a user
func (u *User) generateToken() (string, error) {

	exp := time.Now().Add(2 * time.Hour).Unix()

	claims := &CustomClaims{
		Username: u.Username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: exp,
			Issuer:    "sleazy_e",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenStr, err := token.SignedString("verysecretkey")

	if err != nil {
		return "", err
	}

	return tokenStr, nil
}

// main function
func main() {
	r := gin.Default()

	r.GET("/", indexHandler)
	r.POST("/register", Register)
	r.POST("/login", Login)

	r.Run(":8080")
}
