package composes

import (
	controllers "go_clean_api/api/adapters/controllers"
	presenters "go_clean_api/api/adapters/presenters"
	handler "go_clean_api/api/infra/impl/handlers"
	repo "go_clean_api/api/infra/impl/repositories"
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

	urepo := repo.BuildUserRepoImpl()
	srepo := repo.BuildSessionRepository()
	encrypter := handler.EncrypterHandlerImpl{}
	tmhandler := handler.TokenManagerHandlerImpl{}
	ptr := presenters.LoginPresenter{}
	interactor := interactor.BuildLoginInteractor(urepo, srepo, &encrypter, &tmhandler, &ptr)
	ctrl := controllers.LoginController{Interactor: *interactor}

	output := ctrl.Run(input)

	c.JSON(int(output.StatusCode), output)
}
