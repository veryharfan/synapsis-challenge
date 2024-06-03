package middleware

import (
	"net/http"
	"strings"
	"synapsis-challenge/app/contract"
	"synapsis-challenge/config"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func Authentication(jwtConfig config.JWT) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get token from Authorization header
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, contract.APIResponseErr(contract.ErrUnauthorized))
			return
		}

		// Extract token from header
		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, contract.APIResponseErr(contract.ErrUnauthorized))
			return
		}
		tokenString := parts[1]

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, jwt.ErrSignatureInvalid
			}
			return []byte(jwtConfig.SigningKey), nil
		})
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, contract.APIResponseErr(contract.ErrUnauthorized))
			return
		}

		if !token.Valid {
			c.AbortWithStatusJSON(http.StatusUnauthorized, contract.APIResponseErr(contract.ErrUnauthorized))
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			c.AbortWithStatusJSON(http.StatusUnauthorized, contract.APIResponseErr(contract.ErrUnauthorized))
			return
		}

		exp, ok := claims["exp"].(float64)
		if !ok {
			c.AbortWithStatusJSON(http.StatusUnauthorized, contract.APIResponseErr(contract.ErrUnauthorized))
			return
		}

		expirationTime := time.Unix(int64(exp), 0)
		if time.Now().After(expirationTime) {
			c.AbortWithStatusJSON(http.StatusUnauthorized, contract.APIResponseErr(contract.ErrUnauthorized))
			return
		}

		uid, ok := claims["uid"].(float64)
		if !ok {
			c.AbortWithStatusJSON(http.StatusUnauthorized, contract.APIResponseErr(contract.ErrUnauthorized))
			return
		}

		c.Set("uid", uid)

		c.Next()
	}
}
