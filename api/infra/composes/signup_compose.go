package composes

import (
	controllers "go_clean_api/api/adapters/controllers"
	presenters "go_clean_api/api/adapters/presenters"
	"go_clean_api/api/infra/factories"
	interactor "go_clean_api/api/usecases/signup"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SignUpCompose(c *gin.Context) {

	var input interactor.SignUpInputDTO

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid json"})
		return
	}

	gateway := factories.BuildSignUpGatewayFactory()
	ptr := presenters.SignUpPresenter{}
	interactor := interactor.BuildSignUpInteractor(gateway, &ptr)
	ctrl := controllers.SignUpController{Interactor: *interactor}

	output := ctrl.Run(input)

	c.JSON(int(output.StatusCode), output)
}
