package entities

import "time"

func NewBasicTicketClient() TicketClient {
	return &basicTicketClient{
		creationTime: time.Now(),
		tickets:      make([]Ticket, 0),
	}
}

type TicketClient interface {
	CreatedAt() time.Time
	CreateTicket(title string, description string) error
	TicketCount() int
	GetTickets() []Ticket
}

type basicTicketClient struct {
	creationTime time.Time
	tickets      []Ticket
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
