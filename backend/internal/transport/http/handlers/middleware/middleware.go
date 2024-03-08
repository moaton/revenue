package middleware

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

const (
	authorizationHeader = "Authorization"
)

func Auth() gin.HandlerFunc {
	fn := func(c *gin.Context) {
		header := c.GetHeader(authorizationHeader)
		if header == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, "empty auth header")
			return
		}
		headerParts := strings.Split(header, " ")
		if len(headerParts) != 2 || headerParts[0] != "Bearer" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, "invalid auth header")
			return
		}
		if len(headerParts[1]) == 0 {
			c.AbortWithStatusJSON(http.StatusUnauthorized, "token not found")
			return
		}
		token, err := jwt.Parse(headerParts[1], func(t *jwt.Token) (interface{}, error) {
			if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected sign method: %v", t.Header["alg"])
			}
			return []byte("#3ad$tasd"), nil
		})
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, "token not valid")
			return
		}
		if claims, ok := token.Claims.(jwt.MapClaims); ok {
			if float64(time.Now().UTC().Unix()) > claims["exp"].(float64) {
				c.AbortWithStatusJSON(http.StatusUnauthorized, "token is expired")
				return
			}
			c.Set("userId", claims["userId"].(string))
			c.Next()
		} else {
			c.AbortWithStatusJSON(http.StatusUnauthorized, "token is illegal")
		}
	}
	return gin.HandlerFunc(fn)
}
