package mixins

import (
	handlers_impl "go_clean_api/api/infra/impl/handlers"
	repositories_impl "go_clean_api/api/infra/impl/repositories"
)

type LoginMixin struct {
	*repositories_impl.UserRepositoryImpl
	*repositories_impl.SessionRepositoryImpl
	handlers_impl.EncrypterHandlerImpl
	handlers_impl.TokenManagerHandlerImpl
}

func BuildLoginMixin() *LoginMixin {
	return &LoginMixin{
		UserRepositoryImpl:      repositories_impl.BuildUserRepositoryImpl(),
		SessionRepositoryImpl:   repositories_impl.BuildSessionRepositoryImpl(),
		EncrypterHandlerImpl:    handlers_impl.EncrypterHandlerImpl{},
		TokenManagerHandlerImpl: handlers_impl.TokenManagerHandlerImpl{},
	}
}
