package composes

import (
	controllers "go_clean_api/api/adapters/controllers"
	presenters "go_clean_api/api/adapters/presenters"
	"go_clean_api/api/infra/factories"
	interactor "go_clean_api/api/usecases/login"
	"net/http"

	"github.com/gin-gonic/gin"
)

func LoginCompose(c *gin.Context) {

	var input interactor.LoginInputDTO

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid json"})
		return
	}

	gateway := factories.BuildLoginGatewayFactory()
	ptr := presenters.LoginPresenter{}
	interactor := interactor.BuildLoginInteractor(gateway, &ptr)
	ctrl := controllers.LoginController{Interactor: *interactor}

	output := ctrl.Run(input)

	c.JSON(int(output.StatusCode), output)
}
