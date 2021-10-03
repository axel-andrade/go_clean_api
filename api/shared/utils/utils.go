package utils

import uuid "github.com/satori/go.uuid"

func GenerateUUIDV4() string {
	return uuid.NewV4().String()
}
