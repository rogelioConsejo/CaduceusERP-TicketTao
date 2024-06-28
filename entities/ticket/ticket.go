// Package ticket contains the Ticket interface and a basicTicket struct that implements the Ticket interface.
// The Ticket interface represents a ticket, and the basicTicket struct is a basic implementation of the Ticket interface.
package ticket

import (
	"github.com/google/uuid"
	"ticketTao/entities"
	"time"
)

// Ticket represents an interface for a ticket.
type Ticket interface {
	entities.CreatedEntity
	entities.IdentifiableEntity
	Title() string
	Description() string
	Status() Status
	// AddResponse adds a response to the ticket.
	AddResponse(Response)
	Responses() []Response
	Close()
}

// NewBasicTicket creates a new basic ticket with the given title and description (it has a new ID and creation time
// and is open by default).
func NewBasicTicket(title, description string) Ticket {
	return &basicTicket{
		id:           uuid.New(),
		creationTime: time.Now(),
		title:        title,
		description:  description,
		status:       Open,
	}
}

// MakeEmptyBasicTicket creates a new basic ticket with the given ID, title, and description.
func MakeEmptyBasicTicket(id uuid.UUID, title, description string) Ticket {
	return &basicTicket{
		id:           id,
		creationTime: time.Now(),
		title:        title,
		description:  description,
		status:       Open,
	}
}

// TODO: Function to instantiate a ticket from a database record (an existing ticket).

// Status represents the status of a ticket.
// TODO: Add methods for status transitions.
type Status string

const Open Status = "Open"
const InProgress Status = "InProgress"
const Closed Status = "Closed"

type basicTicket struct {
	creationTime time.Time
	title        string
	description  string
	status       Status
	id           uuid.UUID
	responses    []Response
}

func (b *basicTicket) Close() {
	b.status = Closed
}

func (b *basicTicket) AddResponse(response Response) {
	b.responses = append(b.responses, response)
	b.status = InProgress
}

func (b *basicTicket) Responses() []Response {
	return b.responses
}

func (b *basicTicket) ID() uuid.UUID {
	return b.id
}

func (b *basicTicket) Status() Status {
	return b.status
}

func (b *basicTicket) Title() string {
	return b.title
}

func (b *basicTicket) Description() string {
	return b.description
}

func (b *basicTicket) CreatedAt() time.Time {
	return b.creationTime
}
