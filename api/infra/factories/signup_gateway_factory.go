package factories

import (
	handlers_impl "go_clean_api/api/infra/impl/handlers"
	repositories_impl "go_clean_api/api/infra/impl/repositories"
)

type SignUpGatewayFactory struct {
	*repositories_impl.UserRepositoryImpl
	handlers_impl.EncrypterHandlerImpl
}

func BuildSignUpGatewayFactory() *SignUpGatewayFactory {
	return &SignUpGatewayFactory{
		UserRepositoryImpl:   repositories_impl.BuildUserRepositoryImpl(),
		EncrypterHandlerImpl: handlers_impl.EncrypterHandlerImpl{},
	}
}
