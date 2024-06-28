package client

import (
	"fmt"
	"github.com/google/uuid"
	"ticketTao/entities"
	"ticketTao/entities/ticket"
	"time"
)

// NewBasicTicketClient deprecated, use NewClientFactory to create a client instead
func NewBasicTicketClient(repository ticket.RepositoryClientAccess) TicketClient {
	return &basicTicketClient{
		creationTime:     time.Now(),
		ticketRepository: repository,
		id:               uuid.New(),
	}
}

// InstantiateBasicTicketClient deprecated, use NewClientFactory to instantiate a client instead
func InstantiateBasicTicketClient(id uuid.UUID, ct time.Time, repository ticket.RepositoryClientAccess) TicketClient {
	return &basicTicketClient{
		creationTime:     ct,
		ticketRepository: repository,
		id:               id,
	}
}

// TicketClient represents an interface for a client entity. A client, as a user, is someone who requests
// help (or value, like an order) through the ticket system.
type TicketClient interface {
	entities.IdentifiableEntity
	entities.CreatedEntity
	TicketUser
}

type TicketUser interface {
	TicketClientReader
	TicketWriter
}

type TicketClientReader interface {
	TicketCount() (int, error)
	GetTickets() ([]ticket.Ticket, error)
	GetTicket(uuid uuid.UUID) (ticket.Ticket, error)
}

type TicketWriter interface {
	CreateTicket(title string, description string) error
	AddComment(ticketId uuid.UUID, comment string) error
}

type basicTicketClient struct {
	creationTime     time.Time
	id               uuid.UUID
	ticketRepository ticket.RepositoryClientAccess
}

func (c *basicTicketClient) AddComment(ticketId uuid.UUID, comment string) error {
	tck, err := c.GetTicket(ticketId)
	if err != nil {
		return fmt.Errorf("could not get ticket to add comment: %w", err)
	}
	tck.AddResponse(ticket.NewResponse(c.id, comment))
	err = c.ticketRepository.UpdateTicketForClient(c.id, tck)
	if err != nil {
		return fmt.Errorf("could not update ticket with comment: %w", err)
	}
	return nil
}

func (c *basicTicketClient) ID() uuid.UUID {
	return c.id
}

func (c *basicTicketClient) GetTicket(ticketId uuid.UUID) (ticket.Ticket, error) {
	tick, err := c.ticketRepository.GetTicket(c.id, ticketId)
	if err != nil {
		return nil, fmt.Errorf("could not get ticket with id %s: %w", ticketId.String(), err)
	}
	return tick, nil
}

func (c *basicTicketClient) GetTickets() ([]ticket.Ticket, error) {
	ticks, err := c.ticketRepository.GetAllClientTickets(c.id)
	if err != nil {
		return nil, fmt.Errorf("could not get all tickets: %w", err)
	}
	return ticks, nil
}

func (c *basicTicketClient) TicketCount() (int, error) {
	count, err := c.ticketRepository.GetClientTicketCount(c.id)
	if err != nil {
		return 0, fmt.Errorf("could not get ticket count: %w", err)
	}
	return count, nil
}

func (c *basicTicketClient) CreateTicket(title string, description string) error {
	newTicket, err := ticket.NewBasicTicket(title, description)
	if err != nil {
		return fmt.Errorf("could not create ticket: %w", err)
	}
	err = c.ticketRepository.SaveNewTicketForClient(c.id, newTicket)
	if err != nil {
		return fmt.Errorf("could not create ticket: %w", err)
	}
	return nil
}

func (c *basicTicketClient) CreatedAt() time.Time {
	return c.creationTime
}

func (c *basicTicketClient) CloseTicket(u uuid.UUID) error {
	tck, err := c.GetTicket(u)
	if err != nil {
		return fmt.Errorf("could retrieve ticket to be closed: %w", err)
	}
	tck.Close()
	err = c.ticketRepository.UpdateTicketForClient(c.id, tck)
	if err != nil {
		return fmt.Errorf("could not update ticket: %w", err)
	}
	return nil
}
