package agent

import (
	"errors"
	"fmt"
	"github.com/google/uuid"
	"ticketTao/entities"
	"ticketTao/entities/ticket"
	"time"
)

// New (deprecated, use the Factory methods instead) creates a new instance of Agent.
// It returns an Agent with a randomly generated UUID for the ID
// and the current time for the creation time.
func New(tr ticket.RepositoryAgentAccess) (Agent, error) {
	if tr == nil {
		return nil, ErrTicketRepositoryNotImplemented
	}
	return basicAgent{
		id:               uuid.New(),
		creationTime:     time.Now(),
		ticketRepository: tr,
	}, nil
}

// InstanceAgent (deprecated, use the Factory methods instead) is a function that creates an instance of the Agent
// interface, for an existing Agent.
func InstanceAgent(uuid1 uuid.UUID, creationTime time.Time, tr ticket.RepositoryAgentAccess) (Agent, error) {
	if tr == nil {
		return nil, ErrTicketRepositoryNotImplemented
	}
	return basicAgent{
		id:               uuid1,
		creationTime:     creationTime,
		ticketRepository: tr,
	}, nil
}

// Agent represents an interface for an agent entity. An agent, as a user, is someone who responds to the client's
// requests through the ticket system.
type Agent interface {
	entities.IdentifiableEntity
	entities.CreatedEntity
	// GetTicket retrieves a ticket from the ticket repository based on the provided UUID.
	// If the provided UUID is nil, it returns an error.
	// If the repository returns an error, it returns an error.
	GetTicket(uuid.UUID) (ticket.Ticket, error)
	AnswerTicket(uuid.UUID, string) error
}

type basicAgent struct {
	id               uuid.UUID
	creationTime     time.Time
	ticketRepository ticket.RepositoryAgentAccess
}

func (b basicAgent) AnswerTicket(ticketID uuid.UUID, s string) error {
	commentError := validateTicketComment(ticketID, s)
	if commentError != nil {
		return fmt.Errorf("error while validating comment: %w", commentError)
	}
	tck, err := b.GetTicket(ticketID)
	if err != nil {
		return fmt.Errorf("error while getting ticket: %w", err)
	}
	tck.AddResponse(ticket.NewResponse(b.ID(), s))
	err = b.ticketRepository.UpdateTicket(tck)
	return nil
}

func validateTicketComment(u uuid.UUID, s string) error {
	if u == uuid.Nil {
		return ErrNilTicketID
	}
	if s == "" {
		return errors.New("response cannot be empty")
	}
	return nil
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
