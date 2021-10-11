package factories

import (
	handlers_impl "go_clean_api/api/infra/impl/handlers"
	repositories_impl "go_clean_api/api/infra/impl/repositories"
)

type LoginGatewayFactory struct {
	*repositories_impl.UserRepositoryImpl
	*repositories_impl.SessionRepositoryImpl
	handlers_impl.EncrypterHandlerImpl
	handlers_impl.TokenManagerHandlerImpl
}

func BuildLoginGatewayFactory() *LoginGatewayFactory {
	return &LoginGatewayFactory{
		UserRepositoryImpl:      repositories_impl.BuildUserRepositoryImpl(),
		SessionRepositoryImpl:   repositories_impl.BuildSessionRepositoryImpl(),
		EncrypterHandlerImpl:    handlers_impl.EncrypterHandlerImpl{},
		TokenManagerHandlerImpl: handlers_impl.TokenManagerHandlerImpl{},
	}
}
