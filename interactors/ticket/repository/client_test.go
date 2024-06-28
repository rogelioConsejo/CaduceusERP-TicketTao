package repository

import (
	"errors"
	"github.com/google/uuid"
	"testing"
	"ticketTao/entities/ticket"
)

func TestGetClientTicketRepository(t *testing.T) {
	t.Parallel()
	t.Run("It should return the ticket repository", func(t *testing.T) {
		var clientRepo ticket.RepositoryClientAccess = GetClientTicketRepository()
		if clientRepo == nil {
			t.Error("Repository should not be nil")
		}
	})
}

func TestBasicClientTicketRepository_SaveNewTicketForClient(t *testing.T) {
	t.Parallel()
	var clientRepo ticket.RepositoryClientAccess = GetClientTicketRepository()
	t.Run("It should save a new ticket for a client", func(t *testing.T) {
		var tck ticket.Ticket
		var err error
		tck, err = ticket.NewBasicTicket("title", "description")
		if err != nil {
			t.Error("Error creating new basic ticket")
		}
		var userId uuid.UUID = uuid.New()
		err = clientRepo.SaveNewTicketForClient(userId, tck)
		if err != nil {
			t.Error("Error saving new ticket for client")
		}
	})
	t.Run("It should return an error when the user id is nil", func(t *testing.T) {
		var tck ticket.Ticket
		err := clientRepo.SaveNewTicketForClient(uuid.Nil, tck)
		assertSaveNewTicketForClientError(t, err, ticket.ErrNilCreatorUserID)
	})
	t.Run("It should return an error when the ticket is nil", func(t *testing.T) {
		err := clientRepo.SaveNewTicketForClient(uuid.New(), nil)
		assertSaveNewTicketForClientError(t, err, ticket.ErrNilTicket)
	})
	t.Run("It should return an error when the ticket's Title is empty", func(t *testing.T) {
		t.Skip("This cannot be tested because we cannot create a ticket with an empty title")
		var tck ticket.Ticket
		tck, _ = ticket.NewBasicTicket("", "description")
		err := clientRepo.SaveNewTicketForClient(uuid.New(), tck)
		assertSaveNewTicketForClientError(t, err, ticket.ErrEmptyTitle)
	})
}

func assertSaveNewTicketForClientError(t *testing.T, err error, expected error) {
	t.Helper()
	assertError(t, err, SaveNewTicketForClientError)
	assertError(t, err, expected)
}

func assertError(t *testing.T, err error, expected error) {
	t.Helper()
	if err == nil {
		t.Error("Error should not be nil")
	}
	if !errors.Is(err, expected) {
		t.Errorf("Error should be %v, but is %s", expected, err.Error())
	}
}
