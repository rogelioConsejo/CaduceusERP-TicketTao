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
		var clientRepo ticket.RepositoryClientAccess
		var err error
		clientRepo, err = GetClientTicketRepository(&spyTicketPersistence{})
		if err != nil {
			t.Errorf("Error should be nil, but is %s", err.Error())
		}
		if clientRepo == nil {
			t.Error("Repository should not be nil")
		}
	})
	t.Run("It should return an error when the persistence driver is nil", func(t *testing.T) {
		_, err := GetClientTicketRepository(nil)
		assertErrors(t, err, GetClientTicketRepositoryError, NilPersistenceDriverError)
	})
}

func TestBasicClientTicketRepository_SaveNewTicketForClient(t *testing.T) {
	t.Parallel()
	spyPersistence := &spyTicketPersistence{}
	var tp TicketPersistence = spyPersistence
	var clientRepo ticket.RepositoryClientAccess
	clientRepo, _ = GetClientTicketRepository(tp)
	t.Run("It should save a new ticket for a client", func(t *testing.T) {
		var tck ticket.Ticket
		var err error
		tck, err = ticket.NewBasicTicket("title", "description")
		if err != nil {
			t.Error("Error creating new basic ticket")
		}
		var userId uuid.UUID = uuid.New()
		err = clientRepo.CreateNewTicketForClient(userId, tck)
		if err != nil {
			t.Error("Error saving new ticket for client")
		}
		spyPersistence.assertNewTicketWasSaved(t, userId, tck)
	})
	t.Run("It should return an error when the user id is nil", func(t *testing.T) {
		var tck ticket.Ticket
		err := clientRepo.CreateNewTicketForClient(uuid.Nil, tck)
		assertSaveNewTicketForClientError(t, err, ticket.ErrNilCreatorUserID)
	})
	t.Run("It should return an error when the ticket is nil", func(t *testing.T) {
		err := clientRepo.CreateNewTicketForClient(uuid.New(), nil)
		assertSaveNewTicketForClientError(t, err, ticket.ErrNilTicket)
	})
	t.Run("It should return an error when the ticket's Title is empty", func(t *testing.T) {
		t.Skip("This cannot be tested because we cannot create a ticket with an empty title")
		var tck ticket.Ticket
		tck, _ = ticket.NewBasicTicket("", "description")
		err := clientRepo.CreateNewTicketForClient(uuid.New(), tck)
		assertSaveNewTicketForClientError(t, err, ticket.ErrEmptyTitle)
	})
}

func TestBasicClientTicketRepository_GetTicket(t *testing.T) {
	t.Parallel()
	var clientRepo ticket.RepositoryClientAccess
	spyPersistence := &spyTicketPersistence{}
	clientRepo, _ = GetClientTicketRepository(spyPersistence)
	t.Run("It should return a ticket", func(t *testing.T) {
		clientId := uuid.New()
		spyPersistence.setTicketOwnerOverride(clientId)
		ticketId := uuid.New()
		tck, err := clientRepo.GetTicket(clientId, ticketId)
		if err != nil {
			t.Errorf("Error should be nil, but is %s", err.Error())
		}
		if tck == nil {
			t.Error("Ticket should not be nil")
		}
		spyPersistence.assertTicketOwnerWasChecked(t, ticketId)
		spyPersistence.assertTicketWasRetrieved(t, ticketId)
	})
	t.Run("It should return an error when the ticket does not belong to the client", func(t *testing.T) {
		clientId := uuid.New()
		spyPersistence.setTicketOwnerOverride(uuid.New())
		ticketId := uuid.New()
		_, err := clientRepo.GetTicket(clientId, ticketId)
		assertError(t, err, ErrTicketNotAccessible)
	})
}

func assertErrors(t *testing.T, err error, expected ...error) {
	t.Helper()
	for _, e := range expected {
		assertError(t, err, e)
	}
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

type method string
type argument interface{}

type spyTicketPersistence struct {
	calls               map[method][]argument
	ticketOwnerResponse uuid.UUID
}

func (s *spyTicketPersistence) GetTicket(id uuid.UUID) (ticket.Ticket, error) {
	if s.calls == nil {
		s.calls = make(map[method][]argument)
	}
	s.calls["GetTicket"] = []argument{id}
	return ticket.NewBasicTicket("title", "description")
}

func (s *spyTicketPersistence) GetTicketOwner(ticket uuid.UUID) (client uuid.UUID, err error) {
	if s.calls == nil {
		s.calls = make(map[method][]argument)
	}
	s.calls["GetTicketOwner"] = []argument{ticket}
	if s.ticketOwnerResponse != uuid.Nil {
		return s.ticketOwnerResponse, nil
	}
	return uuid.New(), nil
}

func (s *spyTicketPersistence) SaveNewTicketForClient(client uuid.UUID, tck ticket.Ticket) error {
	if s.calls == nil {
		s.calls = make(map[method][]argument)
	}
	s.calls["SaveNewTicketForClient"] = []argument{client, tck}
	return nil
}

func (s *spyTicketPersistence) assertNewTicketWasSaved(t *testing.T, client uuid.UUID, tck ticket.Ticket) {
	t.Helper()
	if s.calls["SaveNewTicketForClient"] == nil {
		t.Fatal("SaveNewTicketForClient was not called")
	}
	if s.calls["SaveNewTicketForClient"][0].(uuid.UUID) != client {
		t.Error("SaveNewTicketForClient was called with the wrong user id")
	}
	if s.calls["SaveNewTicketForClient"][1].(ticket.Ticket) != tck {
		t.Error("SaveNewTicketForClient was called with the wrong ticket")
	}
}

func (s *spyTicketPersistence) assertTicketOwnerWasChecked(t *testing.T, ticketId uuid.UUID) {
	t.Helper()
	if s.calls["GetTicketOwner"] == nil {
		t.Fatal("GetTicketOwner was not called")
	}
	if s.calls["GetTicketOwner"][0].(uuid.UUID) != ticketId {
		t.Error("GetTicketOwner was called with the wrong ticket id")
	}
}

func (s *spyTicketPersistence) assertTicketWasRetrieved(t *testing.T, ticketId uuid.UUID) {
	t.Helper()
	if s.calls["GetTicket"] == nil {
		t.Fatal("GetTicket was not called")
	}
	if s.calls["GetTicket"][0].(uuid.UUID) != ticketId {
		t.Error("GetTicket was called with the wrong ticket id")
	}
}

func (s *spyTicketPersistence) setTicketOwnerOverride(id uuid.UUID) {
	s.ticketOwnerResponse = id
}
