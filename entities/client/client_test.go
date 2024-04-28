package client

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/rogelioConsejo/golibs/helpers"
	"testing"
	"ticketTao/entities/ticket"
	"time"
)

//TODO: create a client factory to handle adding the ticket repository to clients

func TestNewBasicTicketClient(t *testing.T) {
	t.Parallel()
	client := NewBasicTicketClient(makeFakeTicketRepository())

	t.Run("A basic client can be created and has a creation date", func(t *testing.T) {
		t.Parallel()
		checkClientCreationTime(t, client)
	})
	t.Run("A client has an ID", func(t *testing.T) {
		t.Parallel()
		if client.ID() == uuid.Nil {
			t.Errorf("Expected a non-nil ID, got %s", client.ID().String())
		}
	})
}

func TestInstanceBasicTicketClient(t *testing.T) {
	t.Parallel()
	id := uuid.New()
	creationTime := time.Now()
	var client TicketClient = InstantiateBasicTicketClient(id, creationTime, makeFakeTicketRepository())
	if client.ID() != id {
		t.Errorf("Expected id to be %s, got %s", id.String(), client.ID().String())
	}
}

func TestTicketClient(t *testing.T) {
	t.Parallel()
	client := makeBasicClient(t)
	title := helpers.MakeRandomString(10)
	description := helpers.MakeRandomString(100)
	secondTitle := helpers.MakeRandomString(10)
	secondDescription := helpers.MakeRandomString(100)

	t.Run("A Client can create tickets", func(t *testing.T) {
		err := client.CreateTicket(title, description)
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}
		count, _ := client.TicketCount()
		if count != 1 {
			t.Errorf("Expected ticket count to be 1, got %v", count)
		}

		_ = client.CreateTicket(secondTitle, secondDescription)
		count, _ = client.TicketCount()
		if count != 2 {
			t.Errorf("Expected ticket count to be 2, got %v", count)
		}
	})
	t.Run("A client can return all its tickets", func(t *testing.T) {
		var clientTickets []ticket.Ticket
		clientTickets, err := client.GetTickets()
		if err != nil {
			t.Fatalf("Expected no error, got %v", err)
		}
		if clientTickets[0].Title() != title {
			t.Errorf("Expected title to be %v, got %v", title, clientTickets[0].Title())
		}
		if clientTickets[0].Description() != description {
			t.Errorf("Expected description to be %v, got %v", description, clientTickets[0].Description())
		}
	})
	t.Run("A client can return all their tickets ordered by creation order", func(t *testing.T) {
		clientTickets, err := client.GetTickets()
		if err != nil {
			t.Fatalf("Expected no error, got %v", err)
		}

		if title != clientTickets[0].Title() {
			t.Errorf("Expected title to be %v, got %v", title, clientTickets[0].Title())
		}
		if description != clientTickets[0].Description() {
			t.Errorf("Expected description to be %v, got %v", description, clientTickets[0].Description())
		}

		if secondTitle != clientTickets[1].Title() {
			t.Errorf("Expected title to be %v, got %v", secondTitle, clientTickets[1].Title())
		}
		if secondDescription != clientTickets[1].Description() {
			t.Errorf("Expected description to be %v, got %v", secondDescription, clientTickets[1].Description())
		}
	})
	t.Run("A client can return a ticket by its id", func(t *testing.T) {
		clientTickets, err := client.GetTickets()
		if err != nil {
			t.Fatalf("Expected no error, got %v", err)
		}

		if len(clientTickets) == 0 {
			t.Fatalf("Expected at least one ticket, got none")
		}
		ticketId := clientTickets[0].ID()
		tck, err := client.GetTicket(ticketId)
		if err != nil {
			t.Fatalf("Expected no error, got %v", err)
		}
		if tck.Title() != title {
			t.Errorf("Expected title to be %v, got %v", title, tck.Title())
		}
		if tck.Description() != description {
			t.Errorf("Expected description to be %v, got %v", description, tck.Description())
		}
	})
}

func TestBasicTicketClient_CreateTicket(t *testing.T) {
	t.Parallel()
	title := helpers.MakeRandomString(10)
	description := helpers.MakeRandomString(100)
	t.Run("A client can create a ticket and save it as his in the ticket repository", func(t *testing.T) {
		ticketRepository := makeSpyTicketRepository()
		client := basicTicketClient{
			creationTime:     time.Now(),
			id:               uuid.New(),
			ticketRepository: ticketRepository,
		}
		err := client.CreateTicket(title, description)
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}
		if ticketRepository.calls["SaveNewTicketForClient"][0] != client.ID().String() {
			t.Errorf("Expected saveNewTicket to be called")
		}
	})
}

func TestBasicTicketClient_GetTickets(t *testing.T) {
	t.Parallel()
	t.Run("A client can get all their tickets", func(t *testing.T) {
		ticketRepository := makeSpyTicketRepository()
		client := basicTicketClient{
			creationTime:     time.Now(),
			id:               uuid.New(),
			ticketRepository: ticketRepository,
		}
		_, _ = client.GetTickets()
		if ticketRepository.calls == nil {
			t.Errorf("Expected SaveNewTicketForClient to be called")
		}
		if ticketRepository.calls["GetAllClientTickets"][0] != client.ID().String() {
			t.Errorf("Expected GetAllClientTickets to be called with client id")
		}
	})
}

func TestBasicTicketClient_GetTicket(t *testing.T) {
	t.Parallel()
	t.Run("A client can get a ticket by its id", func(t *testing.T) {
		ticketRepository := makeSpyTicketRepository()
		client := basicTicketClient{
			creationTime:     time.Now(),
			id:               uuid.New(),
			ticketRepository: ticketRepository,
		}
		_, _ = client.GetTicket(uuid.New())
		if ticketRepository.calls == nil {
			t.Errorf("Expected GetTicket to be called")
		}
		if ticketRepository.calls["GetTicket"][0] != client.ID().String() {
			t.Errorf("Expected GetTicket to be called with client id")
		}
		if ticketRepository.calls["GetTicket"][1] == uuid.Nil.String() {
			t.Errorf("Expected GetTicket to be called with ticket id")
		}
	})
}

func TestBasicTicketClient_TicketCount(t *testing.T) {
	t.Parallel()
	t.Run("A client can get the count of their tickets", func(t *testing.T) {
		ticketRepository := makeSpyTicketRepository()
		client := basicTicketClient{
			creationTime:     time.Now(),
			id:               uuid.New(),
			ticketRepository: ticketRepository,
		}
		_, _ = client.TicketCount()
		if ticketRepository.calls == nil {
			t.Errorf("Expected GetClientTicketCount to be called")
		}
		clientId := client.ID().String()
		if ticketRepository.calls["GetClientTicketCount"][0] != clientId {
			t.Errorf("Expected GetClientTicketCount to be called with client id")
		}
	})
}

func checkClientCreationTime(t *testing.T, client TicketClient) {
	t.Helper()
	time.Sleep(11 * time.Millisecond)
	if client.CreatedAt().After(time.Now().Add(-10 * time.Millisecond)) {
		t.Errorf("Expected creation date to be in the past, got %v", client.CreatedAt())
	}
}

func makeBasicClient(t *testing.T) TicketClient {
	t.Helper()
	client := NewBasicTicketClient(makeFakeTicketRepository())
	if client == nil {
		t.Errorf("Expected a client, got nil")
	}
	return client
}

func makeFakeTicketRepository() ticket.RepositoryClientAccess {
	return &fakeTicketRepository{}
}

type fakeTicketRepository struct {
	tickets     []ticket.Ticket
	ticketIndex map[uuid.UUID]ticket.Ticket
}

func (r *fakeTicketRepository) SaveNewTicketForClient(userId uuid.UUID, tck ticket.Ticket) error {
	r.tickets = append(r.tickets, tck)
	if r.ticketIndex == nil {
		r.ticketIndex = make(map[uuid.UUID]ticket.Ticket)
	}
	if tck.ID() == uuid.Nil {
		return fmt.Errorf("ticket id is nil")
	}
	r.ticketIndex[tck.ID()] = tck
	return nil
}

func (r *fakeTicketRepository) GetAllClientTickets(_ uuid.UUID) ([]ticket.Ticket, error) {
	return r.tickets, nil
}

func (r *fakeTicketRepository) GetClientTicketCount(_ uuid.UUID) (int, error) {
	return len(r.tickets), nil
}

func (r *fakeTicketRepository) GetTicket(_, ticket uuid.UUID) (ticket.Ticket, error) {
	tck, ok := r.ticketIndex[ticket]
	if !ok {
		return nil, fmt.Errorf("ticket with id %s not found", ticket.String())
	}
	return tck, nil
}

func makeSpyTicketRepository() *spyTicketRepository {
	return &spyTicketRepository{}
}

type calls map[method]params
type method string
type params []string

type spyTicketRepository struct {
	calls calls
}

func (r *spyTicketRepository) GetTicket(client, ticket uuid.UUID) (ticket.Ticket, error) {
	if r.calls == nil {
		r.calls = make(calls)
	}
	r.calls["GetTicket"] = []string{client.String(), ticket.String()}
	return nil, nil
}

func (r *spyTicketRepository) GetAllClientTickets(id uuid.UUID) ([]ticket.Ticket, error) {
	if r.calls == nil {
		r.calls = make(calls)
	}
	r.calls["GetAllClientTickets"] = []string{id.String()}
	return nil, nil
}

func (r *spyTicketRepository) GetClientTicketCount(clientID uuid.UUID) (int, error) {
	if r.calls == nil {
		r.calls = make(calls)
	}
	r.calls["GetClientTicketCount"] = []string{clientID.String()}
	return 0, nil
}

func (r *spyTicketRepository) SaveNewTicketForClient(client uuid.UUID, ticket ticket.Ticket) error {
	if r.calls == nil {
		r.calls = make(calls)
	}
	r.calls["SaveNewTicketForClient"] = []string{client.String(), ticket.ID().String()}
	return nil
}
