package repository

import (
	"errors"
	"github.com/google/uuid"
	"ticketTao/entities/ticket"
)

// GetClientTicketRepository returns a new instance of ticket.RepositoryClientAccess
func GetClientTicketRepository(tp TicketPersistence) (ticket.RepositoryClientAccess, error) {
	if tp == nil {
		return nil, errors.Join(GetClientTicketRepositoryError, NilPersistenceDriverError)
	}
	return basicClientTicketRepository{tp}, nil
}

type basicClientTicketRepository struct {
	persistence TicketPersistence
}

// GetTicket returns a ticket for a client, it returns an error if the client does not own the ticket or if the persistence
// returns an error.
func (b basicClientTicketRepository) GetTicket(clientId, ticketId uuid.UUID) (ticket.Ticket, error) {
	err := b.validateTicketOwnership(clientId, ticketId)
	if err != nil {
		return nil, errors.Join(GetTicketError, err)

	}
	tck, err := b.persistence.GetTicket(ticketId)
	if err != nil {
		return nil, errors.Join(GetTicketError, err)
	}
	return tck, nil
}

func (b basicClientTicketRepository) validateTicketOwnership(client uuid.UUID, tck uuid.UUID) error {
	owner, err := b.persistence.GetTicketOwner(tck)
	if err != nil {
		return errors.Join(ValidateTicketOwnershipError, err)
	}
	if owner != client {
		return errors.Join(GetTicketError, ErrTicketNotAccessible)
	}
	return nil
}

func (b basicClientTicketRepository) GetAllClientTickets(client uuid.UUID) ([]ticket.Ticket, error) {
	//TODO implement me
	panic("implement me")
}

func (b basicClientTicketRepository) GetClientTicketCount(client uuid.UUID) (int, error) {
	//TODO implement me
	panic("implement me")
}

// CreateNewTicketForClient creates a new ticket for a client, it returns an error if the user id is nil or the ticket
// is nil. It also returns an error if the ticket's title is empty or the ticket's description is empty, to make sure
// that the ticket is valid.
func (b basicClientTicketRepository) CreateNewTicketForClient(userId uuid.UUID, tck ticket.Ticket) error {
	if userId == uuid.Nil {
		return errors.Join(SaveNewTicketForClientError, ticket.ErrNilCreatorUserID)
	}
	if tck == nil {
		return errors.Join(SaveNewTicketForClientError, ticket.ErrNilTicket)
	}
	if tck.Title() == "" {
		return errors.Join(SaveNewTicketForClientError, ticket.ErrEmptyTitle)
	}
	err := b.persistence.SaveNewTicketForClient(userId, tck)
	if err != nil {
		return errors.Join(SaveNewTicketForClientError, err)
	}
	return nil
}

func (b basicClientTicketRepository) UpdateTicketForClient(userId uuid.UUID, tck ticket.Ticket) error {
	//TODO implement me
	panic("implement me")
}

var GetClientTicketRepositoryError error = errors.New("error getting client ticket repository")
var SaveNewTicketForClientError error = errors.New("error saving new ticket for client")
var NilPersistenceDriverError error = errors.New("persistence driver cannot be nil")
var GetTicketError error = errors.New("error getting ticket")
var ValidateTicketOwnershipError error = errors.New("error retrieving ticket owner")

var ErrTicketNotAccessible error = errors.New("ticket is not accessible by the client")
