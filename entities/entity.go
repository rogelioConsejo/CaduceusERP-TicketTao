package entities

import (
	"github.com/google/uuid"
	"time"
)

type CreatedEntity interface {
	CreatedAt() time.Time
}

type IdentifiableEntity interface {
	ID() uuid.UUID
}
