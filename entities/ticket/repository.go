package ticket

import (
	"github.com/google/uuid"
)

type Repository interface {
	Get(id uuid.UUID) Ticket
}
