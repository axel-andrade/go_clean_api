package middlewares

import (
	"fmt"
	handlers_impl "go_clean_api/api/infra/impl/handlers"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func AuthorizeJWT() gin.HandlerFunc {
	return func(c *gin.Context) {

		authHeader := c.GetHeader("Authorization")
		encodedToken := authHeader[len("Bearer"):]
		tmi := handlers_impl.TokenManagerHandlerImpl{}

		token, err := tmi.VerifyToken(encodedToken)
		if token.Valid {
			claims := token.Claims.(jwt.MapClaims)
			fmt.Println(claims)
			return
		}

		fmt.Println(err)
		c.AbortWithStatus(http.StatusUnauthorized)
	}
}
