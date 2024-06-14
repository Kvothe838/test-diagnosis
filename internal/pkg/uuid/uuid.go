package uuid

import (
	"github.com/google/uuid"
)

type UUID interface {
	GetNew() string
}

func NewReal() UUID {
	return real{}
}

type real struct{}

func (real) GetNew() string { return uuid.New().String() }

func NewFake(UUID string) fake {
	return fake{UUID: UUID}
}

type fake struct {
	UUID string
}

func (f fake) GetNew() string {
	return f.UUID
}
