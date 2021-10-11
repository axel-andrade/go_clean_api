package factories

import (
	handlers_impl "go_clean_api/api/infra/impl/handlers"
	repositories_impl "go_clean_api/api/infra/impl/repositories"
)

type LogoutGatewayFactory struct {
	*repositories_impl.UserRepositoryImpl
	*repositories_impl.SessionRepositoryImpl
	handlers_impl.EncrypterHandlerImpl
	handlers_impl.TokenManagerHandlerImpl
}

func BuildLogoutGatewayFactory() *LogoutGatewayFactory {
	return &LogoutGatewayFactory{
		SessionRepositoryImpl:   repositories_impl.BuildSessionRepositoryImpl(),
		TokenManagerHandlerImpl: handlers_impl.TokenManagerHandlerImpl{},
	}
}
