package composes

import (
	controllers "go_clean_api/api/adapters/controllers"
	presenters "go_clean_api/api/adapters/presenters"
	handler "go_clean_api/api/infra/impl/handlers"
	repo "go_clean_api/api/infra/impl/repositories"
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

	repo := repo.BuildUserRepoImpl()
	encrypter := handler.EncrypterHandlerImpl{}
	ptr := presenters.SignUpPresenter{}
	interactor := interactor.BuildSignUpInteractor(repo, &encrypter, &ptr)
	ctrl := controllers.SignUpController{Interactor: *interactor}

	output := ctrl.Run(input)

	c.JSON(int(output.StatusCode), output)
}
