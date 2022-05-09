package middleware

import (
	"errors"
	"server-service/auth"
	"server-service/utils"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func Middleware(authService auth.AuthService) gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" || len(tokenString) == 0 {
			c.AbortWithStatusJSON(401, utils.ErrorMessages(utils.ErrorUnauthorizeUser, errors.New("error user unauthorize")))
			return
		}

		token, err := authService.ValidateToken(tokenString)
		if err != nil {
			c.AbortWithStatusJSON(401, utils.ErrorMessages(utils.ErrorUnauthorizeUser, errors.New("error user unauthorize")))
			return
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			c.Set("role", claims["role"])
			c.Set("user_id", claims["user_id"])

		} else {
			c.AbortWithStatusJSON(401, utils.ErrorMessages(utils.ErrorUnauthorizeUser, errors.New("error user unauthorize")))
			return
		}

		c.Next()
	}
}
