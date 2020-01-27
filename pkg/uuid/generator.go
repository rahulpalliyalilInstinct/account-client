package uuid

import (
	"github.com/satori/go.uuid"
)

type uniqueIdentifier struct{}

func NewUniqueIdentifier() *uniqueIdentifier {
	return &uniqueIdentifier{}
}

func (u *uniqueIdentifier) Generator() string {
	return uuid.NewV4().String()
}
