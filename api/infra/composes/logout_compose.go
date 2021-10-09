package composes

import (
	handlers_impl "go_clean_api/api/infra/impl/handlers"
	repositories_impl "go_clean_api/api/infra/impl/repositories"
	"net/http"

	"github.com/gin-gonic/gin"
)

func LogoutCompose(c *gin.Context) {

	authHeader := c.GetHeader("Authorization")
	encodedToken := authHeader[len("Bearer "):]

	tokenManager := handlers_impl.TokenManagerHandlerImpl{}
	sessionRepo := repositories_impl.BuildSessionRepositoryImpl()

	au, err := tokenManager.ExtractTokenMetadata(encodedToken)
	if err != nil {
		c.JSON(http.StatusUnauthorized, "unauthorized")
		return
	}

	deleted, err := sessionRepo.DeleteAuth(au.AccessUUID)
	if err != nil || deleted == 0 { //if any goes wrong
		c.JSON(http.StatusUnauthorized, "unauthorized")
		return
	}

	c.JSON(http.StatusOK, "Successfully logged out")
}
