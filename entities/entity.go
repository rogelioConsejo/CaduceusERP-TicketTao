package entities

import (
	"github.com/google/uuid"
	"testing"
	"time"
)

type CreatedEntity interface {
	CreatedAt() time.Time
}

type IdentifiableEntity interface {
	ID() uuid.UUID
}

func TestIdentifiableEntity(t *testing.T, entity IdentifiableEntity) {
	if entity.ID().String() == "" {
		t.Error("ID should not be empty")
	}
	var _ uuid.UUID = entity.ID()
}
