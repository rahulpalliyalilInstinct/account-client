package uuid

import (
	"github.com/satori/go.uuid"
)

type uniqueIdentifier struct{}

// NewUniqueIdentifier returns a uniqueIdentifier
func NewUniqueIdentifier() *uniqueIdentifier {
	return &uniqueIdentifier{}
}

// NewV4 returns random generated canonical string representation of UUID.
func (u *uniqueIdentifier) Generator() string {
	return uuid.NewV4().String()
}
