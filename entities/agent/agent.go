package agent

import (
	"errors"
	"fmt"
	"github.com/google/uuid"
	"ticketTao/entities"
	"ticketTao/entities/ticket"
	"time"
)

// TODO: Create an AgentFactory that handles the ticket repository while creating or instantiating an Agent to avoid the need to pass the ticket repository every time.
// TODO: Create an TicketRepositoryAgentAccess interface to handle the ticket repository for agents.

// New creates a new instance of Agent.
// It takes a ticket repository as input parameter.
// It returns an Agent with a randomly generated UUID for the ID
// and the current time for the creation time.
func New(tr ticket.RepositoryClientAccess) (Agent, error) {
	if tr == nil {
		return nil, ErrTicketRepositoryNotImplemented
	}
	return basicAgent{
		id:               uuid.New(),
		creationTime:     time.Now(),
		ticketRepository: tr,
	}, nil
}

// InstanceAgent is a function that creates an instance of the Agent interface.
// It takes an uuid.UUID (uuid1) and a time.Time (creationTime) as input parameters, as well as a ticket repository.
// and returns an Agent instance.
func InstanceAgent(uuid1 uuid.UUID, creationTime time.Time, tr ticket.RepositoryClientAccess) (Agent, error) {
	if tr == nil {
		return nil, ErrTicketRepositoryNotImplemented
	}
	return basicAgent{
		id:               uuid1,
		creationTime:     creationTime,
		ticketRepository: tr,
	}, nil
}

// Agent represents an interface for an agent entity.
// It extends the entities.IdentifiableEntity and entities.CreatedEntity interfaces.
// An agent can retrieve a ticket by its UUID using the GetTicket method.
type Agent interface {
	entities.IdentifiableEntity
	entities.CreatedEntity
	// GetTicket retrieves a ticket from the ticket repository based on the provided UUID.
	// It returns the ticket retrieved from the repository and any error encountered during the process.
	// If the provided UUID is nil, it returns an error.
	// If the repository returns an error, it returns an error.
	GetTicket(uuid.UUID) (ticket.Ticket, error)
}

type basicAgent struct {
	id               uuid.UUID
	creationTime     time.Time
	ticketRepository ticket.RepositoryClientAccess
}

func (b basicAgent) CreatedAt() time.Time {
	return b.creationTime
}

func (b basicAgent) GetTicket(id uuid.UUID) (ticket.Ticket, error) {
	if id == uuid.Nil {
		return nil, ErrNilTicketID
	}
	tck, err := b.ticketRepository.GetTicket(id)
	if err != nil {
		return nil, fmt.Errorf("%w: %w", ErrRetrievingTicket, err)
	}
	return tck, nil
}

func (b basicAgent) ID() uuid.UUID {
	return b.id
}

var ErrTicketRepositoryNotImplemented = errors.New("ticket repository not implemented")
var ErrNilTicketID = errors.New("nil ticket ID")
var ErrRetrievingTicket = errors.New("error while retrieving ticket")
