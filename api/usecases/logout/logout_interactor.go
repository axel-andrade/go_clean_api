package logout

import (
	"fmt"
	ERROR "go_clean_api/api/shared/constants/errors"
	"go_clean_api/api/usecases/common"
)

type LogoutInteractor struct {
	Gateway   LogoutGateway
	Presenter LogoutPresenter
}

func BuildLogoutInteractor(g LogoutGateway, p LogoutPresenter) *LogoutInteractor {
	return &LogoutInteractor{Gateway: g, Presenter: p}
}

func (bs *LogoutInteractor) Execute(encodedToken string) common.OutputPort {

	au, err := bs.Gateway.ExtractTokenMetadata(encodedToken)
	if err != nil {
		return bs.Presenter.Show(fmt.Errorf(ERROR.UNAUTHORIZED))
	}

	deleted, err := bs.Gateway.DeleteAuth(au.AccessUUID)
	if err != nil || deleted == 0 {
		return bs.Presenter.Show(fmt.Errorf(ERROR.UNAUTHORIZED))
	}

	return bs.Presenter.Show(nil)
}
