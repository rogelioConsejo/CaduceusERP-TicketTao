package entities

import (
	"errors"
	"github.com/google/uuid"
	"time"
)

type IdentifiableEntity interface {
	ID() uuid.UUID
}

type CreatedEntity interface {
	CreatedAt() time.Time
}

var ErrNilID error = errors.New("entity ID cannot be nil")
var ErrNilCreationTime error = errors.New("entity creation time cannot be nil")
var ErrFutureCreationTime error = errors.New("entity creation time cannot be in the future")
