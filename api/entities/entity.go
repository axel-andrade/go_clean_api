package entities

import "go_clean_api/api/shared/utils"

// ID entity ID
type EntityID = string

// NewID create a new entity ID
func NewID() EntityID {
	return EntityID(utils.GenerateUUIDV4())
}
