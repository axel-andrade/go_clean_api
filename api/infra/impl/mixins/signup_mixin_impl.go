package mixins

import (
	handlers_impl "go_clean_api/api/infra/impl/handlers"
	repositories_impl "go_clean_api/api/infra/impl/repositories"
)

type SignUpMixin struct {
	*repositories_impl.UserRepositoryImpl
	handlers_impl.EncrypterHandlerImpl
}

func BuildSignUpMixin() *SignUpMixin {
	return &SignUpMixin{
		UserRepositoryImpl:   repositories_impl.BuildUserRepositoryImpl(),
		EncrypterHandlerImpl: handlers_impl.EncrypterHandlerImpl{},
	}
}
