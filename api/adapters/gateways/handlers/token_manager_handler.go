package handlers

import (
	"go_clean_api/api/entities"
)

type TokenManagerHandler interface {
	GenerateToken(userid entities.EntityID) (*entities.TokenDetails, error)
	ExtractTokenMetadata(encoded string) (*entities.AccessDetails, error)
}
