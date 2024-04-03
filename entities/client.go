package entities

import "time"

func NewBasicTicketClient() TicketClient {
	return &basicTicketClient{
		creationTime: time.Now(),
		tickets:      make([]Ticket, 0),
	}
}

type TicketClient interface {
	ClientTicketManager
	CreatedEntity
}

type ClientTicketManager interface {
	ClientTicketReader
	ClientTicketWriter
}

type ClientTicketReader interface {
	TicketCount() int
	GetTickets() []Ticket
	GetTicket(index int) Ticket
}

type ClientTicketWriter interface {
	CreateTicket(title string, description string) error
}

type basicTicketClient struct {
	creationTime time.Time
	tickets      []Ticket
}

func (b *basicTicketClient) GetTicket(index int) Ticket {
	return b.tickets[index]
}

func (b *basicTicketClient) GetTickets() []Ticket {
	return b.tickets
}

func (b *basicTicketClient) TicketCount() int {
	return len(b.tickets)
}

func (b *basicTicketClient) CreateTicket(title string, description string) error {
	b.tickets = append(b.tickets, NewBasicTicket(title, description))
	return nil
}

func (b *basicTicketClient) CreatedAt() time.Time {
	return b.creationTime
}
