package client

import (
	"ticketTao/entities"
	"ticketTao/entities/ticket"
	"time"
)

func NewBasicTicketClient() TicketClient {
	return &basicTicketClient{
		creationTime: time.Now(),
		tickets:      make([]ticket.Ticket, 0),
	}
}

type TicketClient interface {
	ClientTicketManager
	entities.CreatedEntity
}

type ClientTicketManager interface {
	ClientTicketReader
	ClientTicketWriter
}

type ClientTicketReader interface {
	TicketCount() int
	GetTickets() []ticket.Ticket
	GetTicket(index int) ticket.Ticket
}

type ClientTicketWriter interface {
	CreateTicket(title string, description string) error
}

type basicTicketClient struct {
	creationTime time.Time
	tickets      []ticket.Ticket
}

func (b *basicTicketClient) GetTicket(index int) ticket.Ticket {
	return b.tickets[index]
}

func (b *basicTicketClient) GetTickets() []ticket.Ticket {
	return b.tickets
}

func (b *basicTicketClient) TicketCount() int {
	return len(b.tickets)
}

func (b *basicTicketClient) CreateTicket(title string, description string) error {
	b.tickets = append(b.tickets, ticket.NewBasicTicket(title, description))
	return nil
}

func (b *basicTicketClient) CreatedAt() time.Time {
	return b.creationTime
}
