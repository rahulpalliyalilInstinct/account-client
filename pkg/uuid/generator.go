package uuid

import (
	uuid "github.com/satori/go.uuid"
)

type UniqueIdentifier struct{}

// NewUniqueIdentifier returns a UniqueIdentifier
func NewUniqueIdentifier() *UniqueIdentifier {
	return &UniqueIdentifier{}
}

// NewV4 returns random generated canonical string representation of UUID.
func (u *UniqueIdentifier) Generator() string {
	return uuid.NewV4().String()
}
