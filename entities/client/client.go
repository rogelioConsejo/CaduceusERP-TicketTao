package client

import (
	"fmt"
	"github.com/google/uuid"
	"ticketTao/entities"
	"ticketTao/entities/ticket"
	"time"
)

func NewBasicTicketClient(repository ticket.RepositoryClientAccess) TicketClient {
	return &basicTicketClient{
		creationTime:     time.Now(),
		ticketRepository: repository,
		id:               uuid.New(),
	}
}

func InstantiateBasicTicketClient(id uuid.UUID, ct time.Time, repository ticket.RepositoryClientAccess) TicketClient {
	return &basicTicketClient{
		creationTime:     ct,
		ticketRepository: repository,
		id:               id,
	}
}

type TicketClient interface {
	entities.IdentifiableEntity
	TicketUser
	entities.CreatedEntity
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
}

type basicTicketClient struct {
	creationTime     time.Time
	id               uuid.UUID
	ticketRepository ticket.RepositoryClientAccess
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
	newTicket := ticket.NewBasicTicket(title, description)
	err := c.ticketRepository.SaveNewTicketForClient(c.id, newTicket)
	if err != nil {
		return fmt.Errorf("could not create ticket: %w", err)
	}
	return nil
}

func (c *basicTicketClient) CreatedAt() time.Time {
	return c.creationTime
}
