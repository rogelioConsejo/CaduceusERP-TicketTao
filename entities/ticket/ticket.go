// Package ticket contains the Ticket interface and a basicTicket struct that implements the Ticket interface.
// The Ticket interface represents a ticket, and the basicTicket struct is a basic implementation of the Ticket interface.
package ticket

import (
	"errors"
	"github.com/google/uuid"
	"ticketTao/entities"
	"time"
)

// NewBasicTicket creates a new basic ticket with the given title and description (it has a new ID and creation time
// and is open by default). A ticket cannot be created with an empty title.
func NewBasicTicket(title, description string) (Ticket, error) {
	if title == "" {
		return nil, errors.Join(NewBasicTicketError, ErrEmptyTitle)
	}
	return &basicTicket{
		id:           uuid.New(),
		creationTime: time.Now(),
		title:        title,
		description:  description,
		status:       Open,
	}, nil
}

// MakeBasicTicket creates a new basic ticket with the given ID, creation time, and data.
func MakeBasicTicket(id uuid.UUID, creationTime time.Time, data Data) (Ticket, error) {
	if id == uuid.Nil {
		return nil, errors.Join(NewBasicTicketError, entities.ErrNilID)
	}
	if creationTime.IsZero() {
		return nil, errors.Join(NewBasicTicketError, entities.ErrNilCreationTime)

	}
	if creationTime.After(time.Now()) {
		return nil, errors.Join(NewBasicTicketError, entities.ErrFutureCreationTime)
	}
	if data.Title == "" {
		return nil, errors.Join(NewBasicTicketError, ErrEmptyTitle)
	}
	if data.Status == "" {
		return nil, errors.Join(NewBasicTicketError, ErrEmptyStatus)
	}
	return &basicTicket{
		creationTime, data.Title, data.Description, data.Status, id, data.Responses}, nil
}

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

// Data represents the data of a ticket.
type Data struct {
	Title       string     `json:"title"`
	Description string     `json:"description"`
	Status      Status     `json:"status"`
	Responses   []Response `json:"responses"`
}

// Status represents the status of a ticket.
// TODO: Add methods for status transitions.
type Status string

// Open is the status of a ticket when it is created.
const Open Status = "Open"

// InProgress is the status of a ticket when it has responses.
const InProgress Status = "InProgress"

// Closed is the status of a ticket when it is done.
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

var NewBasicTicketError error = errors.New("error creating new basic ticket")
var ErrEmptyTitle error = errors.New("ticket title cannot be empty")
var ErrEmptyStatus error = errors.New("ticket status cannot be empty")
