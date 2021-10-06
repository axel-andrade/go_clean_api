package middlewares

import (
	handlers_impl "go_clean_api/api/infra/impl/handlers"
	repositories_impl "go_clean_api/api/infra/impl/repositories"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Authorize() gin.HandlerFunc {
	return func(c *gin.Context) {

		tokenManager := handlers_impl.TokenManagerHandlerImpl{}
		sessionRepo := repositories_impl.BuildSessionRepository()

		authHeader := c.GetHeader("Authorization")
		encodedToken := authHeader[len("Bearer "):]

		tokenAuth, err := tokenManager.ExtractTokenMetadata(encodedToken)
		if err != nil {
			c.JSON(http.StatusUnauthorized, err.Error())
			c.Abort()
			return
		}

		userId, err := sessionRepo.GetAuth(tokenAuth)
		if err != nil {
			c.JSON(http.StatusUnauthorized, err.Error())
			c.Abort()
			return
		}

		// TODO: verificar se o usuario existe no banco de dados

		c.Set("user-id", userId)

		c.Next()
	}
}
