package logout

import (
	"go_clean_api/api/entities"
	"go_clean_api/api/usecases/common"
)

type LogoutGateway interface {
	ExtractTokenMetadata(encoded string) (*entities.AccessDetails, error)
	DeleteAuth(uuid string) (int64, error)
}

type LogoutPresenter interface {
	Show(err error) common.OutputPort
}
