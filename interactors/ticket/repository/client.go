package repository

import (
	"errors"
	"github.com/google/uuid"
	"ticketTao/entities/ticket"
)

// GetClientTicketRepository returns a new instance of ticket.RepositoryClientAccess
func GetClientTicketRepository() ticket.RepositoryClientAccess {
	return basicClientTicketRepository{}
}

type basicClientTicketRepository struct {
}

func (b basicClientTicketRepository) GetTicket(client, ticket uuid.UUID) (ticket.Ticket, error) {
	//TODO implement me
	panic("implement me")
}

func (b basicClientTicketRepository) GetAllClientTickets(client uuid.UUID) ([]ticket.Ticket, error) {
	//TODO implement me
	panic("implement me")
}

func (b basicClientTicketRepository) GetClientTicketCount(client uuid.UUID) (int, error) {
	//TODO implement me
	panic("implement me")
}

// SaveNewTicketForClient saves a new ticket for a client, it returns an error if the user id is nil or the ticket is nil.
func (b basicClientTicketRepository) SaveNewTicketForClient(userId uuid.UUID, tck ticket.Ticket) error {
	if userId == uuid.Nil {
		return errors.Join(SaveNewTicketForClientError, ticket.ErrNilCreatorUserID)
	}
	if tck == nil {
		return errors.Join(SaveNewTicketForClientError, ticket.ErrNilTicket)
	}
	if tck.Title() == "" {
		return errors.Join(SaveNewTicketForClientError, ticket.ErrEmptyTitle)
	}
	if tck.Description() == "" {
		return errors.Join(SaveNewTicketForClientError, ticket.ErrEmptyDescription)
	}
	return nil
}

func (b basicClientTicketRepository) UpdateTicketForClient(userId uuid.UUID, tck ticket.Ticket) error {
	//TODO implement me
	panic("implement me")
}

var SaveNewTicketForClientError error = errors.New("error saving new ticket for client")
