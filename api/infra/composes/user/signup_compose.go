package composes

import (
	controllers "go_clean_api/api/adapters/controllers/user"
	presenters "go_clean_api/api/adapters/presenters/user"
	"go_clean_api/api/infra/database"
	handlerImpl "go_clean_api/api/infra/impl/handlers"
	repoImpl "go_clean_api/api/infra/impl/repositories"
	interactor "go_clean_api/api/usecases/user"
)

func SignUpCompose() *controllers.SignUpController {
	db := database.GetDB()
	repo := repoImpl.UserRepoImpl{Db: db}
	encrypter := handlerImpl.EncrypterHandlerImpl{}
	prt := presenters.SignUpPresenter{}
	bs := interactor.BuildSignUpInteractor(repo, &encrypter, prt)
	ctrl := controllers.SignUpController{Interactor: *bs}
	return &ctrl
}
