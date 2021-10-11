package composes

import (
	"go_clean_api/api/adapters/controllers"
	"go_clean_api/api/adapters/presenters"
	"go_clean_api/api/infra/factories"
	interactor "go_clean_api/api/usecases/logout"

	"github.com/gin-gonic/gin"
)

func LogoutCompose(c *gin.Context) {

	authHeader := c.GetHeader("Authorization")
	encodedToken := authHeader[len("Bearer "):]

	gateway := factories.BuildLogoutGatewayFactory()
	ptr := presenters.LogoutPresenter{}
	interactor := interactor.BuildLogoutInteractor(gateway, &ptr)
	ctrl := controllers.LogoutController{Interactor: *interactor}

	output := ctrl.Run(encodedToken)

	c.JSON(int(output.StatusCode), output)
}
