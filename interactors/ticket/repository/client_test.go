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

// TODO: We should probably move the ticket title and description validation to the ticket entity.
func TestBasicClientTicketRepository_SaveNewTicketForClient(t *testing.T) {
	t.Parallel()
	var clientRepo ticket.RepositoryClientAccess = GetClientTicketRepository()
	t.Run("It should save a new ticket for a client", func(t *testing.T) {
		var tck ticket.Ticket = ticket.NewBasicTicket("title", "description")
		var userId uuid.UUID = uuid.New()
		err := clientRepo.SaveNewTicketForClient(userId, tck)
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
		var tck ticket.Ticket = ticket.NewBasicTicket("", "description")
		err := clientRepo.SaveNewTicketForClient(uuid.New(), tck)
		assertSaveNewTicketForClientError(t, err, ticket.ErrEmptyTitle)
	})
	t.Run("It should return an error when the ticket's Description is empty", func(t *testing.T) {
		var tck ticket.Ticket = ticket.NewBasicTicket("title", "")
		err := clientRepo.SaveNewTicketForClient(uuid.New(), tck)
		assertSaveNewTicketForClientError(t, err, ticket.ErrEmptyDescription)
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
		t.Errorf("Error should be %v", expected)
	}
}
