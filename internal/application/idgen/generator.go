package idgen

import "github.com/google/uuid"

type DefaultUUIDGenerator struct{}

func (g *DefaultUUIDGenerator) New() string {
	return uuid.New().String()
}
