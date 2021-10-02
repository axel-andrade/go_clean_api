package composes

import (
	controllers "go_clean_api/api/adapters/controllers/user"
	"go_clean_api/api/infra/database"
	repositories_impl "go_clean_api/api/infra/database/repositories"
	usecases "go_clean_api/api/usecases/user"
)

func SignUpCompose() *controllers.SignUpController {
	db := database.GetDB()
	repo := repositories_impl.UserRepoImpl{Db: db}
	bs := usecases.BuildSignUpInteractor(repo)
	ctrl := controllers.SignUpController{Interactor: *bs}
	return &ctrl
}
