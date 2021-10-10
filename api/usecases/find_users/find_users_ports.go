package find_users

import (
	"go_clean_api/api/entities"
	"go_clean_api/api/usecases/common"
)

type FindUsersGateway interface {
	FindUsersPaginate(pagination *entities.PaginationOptions) (*entities.PaginateResult, error)
}

type FindUserInputDTO = entities.PaginationOptions

type FindUsersOutputDTO = entities.PaginateResult

type FindUsersPresenter interface {
	Show(out *FindUsersOutputDTO, err error) common.OutputPort
}
