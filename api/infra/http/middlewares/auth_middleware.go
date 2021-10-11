package middlewares

import (
	"fmt"
	handlers_impl "go_clean_api/api/infra/impl/handlers"
	repositories_impl "go_clean_api/api/infra/impl/repositories"
	ERROR "go_clean_api/api/shared/constants/errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Authorize() gin.HandlerFunc {
	return func(c *gin.Context) {

		tokenManager := handlers_impl.TokenManagerHandlerImpl{}
		sessionRepo := repositories_impl.BuildSessionRepositoryImpl()

		authHeader := c.GetHeader("Authorization")
		encodedToken := authHeader[len("Bearer "):]

		tokenAuth, err := tokenManager.ExtractTokenMetadata(encodedToken)
		if err != nil {
			fmt.Println("error: error in extract token metadata: ", err.Error())
			c.JSON(http.StatusUnauthorized, gin.H{"error": ERROR.UNAUTHORIZED})
			c.Abort()
			return
		}

		userId, err := sessionRepo.GetAuth(tokenAuth)
		if err != nil {
			fmt.Println("error in get auth: ", err.Error())
			c.JSON(http.StatusUnauthorized, gin.H{"error": ERROR.UNAUTHORIZED})
			c.Abort()
			return
		}

		// TODO: verificar se o usuario existe no banco de dados

		c.Set("user-id", userId)

		c.Next()
	}
}
