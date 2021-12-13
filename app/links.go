package pkg

import (
	"github.com/google/uuid"
	"net/url"
)

type Link struct {
	ID uuid.UUID
	count int
	Native url.URL
	Short string
}

type LinkStore interface {
	Create(l Link) (*uuid.UUID, error)
}

