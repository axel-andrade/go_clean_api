package mixins

import (
	repositories_impl "go_clean_api/api/infra/impl/repositories"
)

type FindUsersMixin struct {
	*repositories_impl.UserRepositoryImpl
}

func BuildFindUsersMixin() *FindUsersMixin {
	return &FindUsersMixin{UserRepositoryImpl: repositories_impl.BuildUserRepositoryImpl()}
}
