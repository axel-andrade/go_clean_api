package composes

import (
	controllers "go_clean_api/api/adapters/controllers/user"
	presenters "go_clean_api/api/adapters/presenters/user"
	handlerImpl "go_clean_api/api/infra/impl/handlers"
	repoImpl "go_clean_api/api/infra/impl/repositories"
	interactor "go_clean_api/api/usecases/user/signup"
)

func SignUpCompose() *controllers.SignUpController {
	r := repoImpl.BuildUserRepoImpl()
	e := handlerImpl.EncrypterHandlerImpl{}
	p := presenters.SignUpPresenter{}
	b := interactor.BuildSignUpInteractor(r, &e, &p)
	c := controllers.SignUpController{Interactor: *b}
	return &c
}
