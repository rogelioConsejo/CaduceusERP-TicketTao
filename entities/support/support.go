package support

import (
	"github.com/google/uuid"
	"ticketTao/entities"
	"ticketTao/entities/ticket"
	"time"
)

func NewAgent() Agent {
	return basicAgent{
		id:           uuid.New(),
		creationTime: time.Now(),
	}
}

func InstanceSupportAgent(uuid1 uuid.UUID) Agent {
	return basicAgent{
		id: uuid1,
	}
}

type Agent interface {
	entities.IdentifiableEntity
	entities.CreatedEntity
	GetTicket(testUUID uuid.UUID) ticket.Ticket
}

type basicAgent struct {
	id               uuid.UUID
	creationTime     time.Time
	ticketRepository ticket.Repository
}

func (b basicAgent) CreatedAt() time.Time {
	return b.creationTime
}

func (b basicAgent) GetTicket(id uuid.UUID) ticket.Ticket {
	return b.ticketRepository.Get(id)
}

func (b basicAgent) ID() uuid.UUID {
	return b.id
}
