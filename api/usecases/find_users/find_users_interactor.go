package find_users

import (
	"go_clean_api/api/usecases/common"
)

type FindUsersInteractor struct {
	Gateway   FindUsersGateway
	Presenter FindUsersPresenter
}

func BuildFindUsersInteractor(g FindUsersGateway, p FindUsersPresenter) *FindUsersInteractor {
	return &FindUsersInteractor{Gateway: g, Presenter: p}
}

func (bs *FindUsersInteractor) Execute(input FindUserInputDTO) common.OutputPort {
	data, err := bs.Gateway.FindUsersPaginate(&input)
	if err != nil {
		return bs.Presenter.Show(nil, err)
	}

	return bs.Presenter.Show(data, nil)
}
