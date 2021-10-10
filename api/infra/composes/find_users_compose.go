package composes

import (
	"go_clean_api/api/adapters/controllers"
	"go_clean_api/api/adapters/presenters"
	"go_clean_api/api/infra/impl/mixins"
	"go_clean_api/api/shared/utils"
	interactor "go_clean_api/api/usecases/find_users"
	"net/http"

	"github.com/gin-gonic/gin"
)

func FindUsersCompose(c *gin.Context) {
	gateway := mixins.BuildFindUsersMixin()
	ptr := presenters.FindusersPresenter{}
	interactor := interactor.BuildFindUsersInteractor(gateway, &ptr)
	ctrl := controllers.FindUsersController{Interactor: *interactor}
	output := ctrl.Run(utils.GetPaginationOptionsFromURL(c))
	c.JSON(http.StatusOK, output)
}
