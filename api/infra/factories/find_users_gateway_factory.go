package factories

import (
	repositories_impl "go_clean_api/api/infra/impl/repositories"
)

type FindUsersGatewayFactory struct {
	*repositories_impl.UserRepositoryImpl
}

func BuildFindUsersGatewayFactory() *FindUsersGatewayFactory {
	return &FindUsersGatewayFactory{UserRepositoryImpl: repositories_impl.BuildUserRepositoryImpl()}
}
