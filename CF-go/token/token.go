package token

import (
	model "cfapiapp/models"
	"fmt"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

var secretKey = []byte("secret-key")

func GenerateToken(user model.User) (string, error) {
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["email"] = user.Email
	claims["username"] = user.Username
	claims["codeforces_handle"] = user.Codeforces_Handle
	claims["subscribedblog"] = user.SubscribedBlogsId
	claims["exp"] = time.Now().Add(time.Hour * 24 * 30).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(secretKey)
}

func TokenValid(c *gin.Context) (jwt.MapClaims, error) {
	tokenString := ExtractToken(c)
	// log.Println("TokenString : ", tokenString)
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}
		return secretKey, nil
	})
	if err != nil {
		// log.Println("ErrorTokenValid2 : ", err)
		return nil, err
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, fmt.Errorf("invalid token")
}
func ExtractToken(c *gin.Context) string {
	// Check if the token is present in the query parameter
	token := c.Query("token")
	if token != "" {
		return token
	}

	// Check if the token is present in the Authorization header
	authHeader := c.GetHeader("Authorization")
	if authHeader != "" {
		// Token might be in the format: "Bearer <token>"
		// Split the header value and extract the token part
		parts := strings.Split(authHeader, " ")
		if len(parts) == 2 {
			return parts[1]
		}
	}

	// Check if the token is present in the cookie
	cookie, err := c.Cookie("jwt")
	if err == nil && cookie != "" {
		return cookie
	}

	// If token is not found, return an empty string
	return ""
}
