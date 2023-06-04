package userapi

import (
	"fmt"
	userschema "hoainam/gin-test/user_app/schemas"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	var data userschema.LoginSchema

	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	jwt, err := GenerateJWT(data.Username)

	if err != nil {
		c.JSON(500, err)
	} else {
		c.JSON(http.StatusOK, gin.H{
			"jwt":      jwt,
			"username": data.Username,
		})
	}

}

func GenerateJWT(usernam string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)

	claims["authorized"] = true
	claims["exp"] = time.Now().Add(time.Minute * 30).Unix()

	tokenString, err := token.SignedString("21312312312")

	if err != nil {
		fmt.Errorf("Something Went Wrong: %s", err.Error())
		return "", err
	}

	return tokenString, nil
}
