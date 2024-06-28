package agent

import (
	"errors"
	"fmt"
	"github.com/google/uuid"
	"math/rand"
	"testing"
	"ticketTao/entities/ticket"
	"time"
)

func TestNew(t *testing.T) {
	t.Parallel()
	var agent Agent
	agent, newAgentError := New(stubTicketRepository{})
	if newAgentError != nil {
		t.Fatal("Error should be nil")
	}

	t.Run("An agent should have a non nil UUID", func(t *testing.T) {
		t.Parallel()
		// assert that the ID UUID is not empty
		if agent.ID().String() == "" {
			t.Error("ID should not be empty")
		}
		// assert that the ID method returns a UUID
		var _ uuid.UUID = agent.ID()
		if agent.ID() == uuid.Nil {
			t.Error("ID should not be nil")
		}
	})

	t.Run("An agent should have a creation time", func(t *testing.T) {
		t.Parallel()
		if agent.CreatedAt().IsZero() {
			t.Error("Creation time should not be zero")
		}
		if agent.CreatedAt().After(time.Now()) {
			t.Error("Creation time should be before now")
		}
	})

	t.Run("New should return a ErrorTicketRepositoryNotImplemented if the ticket repository is nil", func(t *testing.T) {
		t.Parallel()
		_, err := New(nil)
		if err == nil {
			t.Error("Error should not be nil")
		}
		if !errors.Is(err, ErrTicketRepositoryNotImplemented) {
			t.Errorf("Error should be %v", ErrTicketRepositoryNotImplemented)
		}
	})
}

func TestInstanceAgent(t *testing.T) {
	t.Parallel()
	var uuid1 uuid.UUID = uuid.New()
	var duration time.Duration = time.Duration(rand.Intn(1000)) * time.Millisecond
	var creationTime time.Time = time.Now().Add(duration)
	var agent Agent
	agent, err := InstanceAgent(uuid1, creationTime, stubTicketRepository{})
	if err != nil {
		t.Fatal("Error should be nil")
	}
	if agent.ID() != uuid1 {
		t.Error("ID should be equal to the one used to create the agent")
	}
	if agent.CreatedAt() != creationTime {
		t.Error("Creation time should be equal to the one used to create the agent")
	}
	t.Run("Should return an error if the ticket repository is nil", func(t *testing.T) {
		t.Parallel()
		_, nilRepoError := InstanceAgent(uuid1, creationTime, nil)
		if nilRepoError == nil {
			t.Error("Error should not be nil")
		}
		if !errors.Is(nilRepoError, ErrTicketRepositoryNotImplemented) {
			t.Errorf("Error should be %v", ErrTicketRepositoryNotImplemented)
		}
	})
}

func TestBasicAgent_GetTicket(t *testing.T) {
	t.Parallel()
	agent, instanceError := InstanceAgent(uuid.New(), time.Now(), stubTicketRepository{})
	if instanceError != nil {
		t.Fatal("Error should be nil")
	}
	t.Run("Should return a ticket from the ticket repository", func(t *testing.T) {
		t.Parallel()
		testUUID := uuid.New()
		tkt, err := agent.GetTicket(testUUID)
		if err != nil {
			t.Fatal("Error should be nil")
		}
		if tkt.ID() != testUUID {
			t.Error("Ticket ID should be equal to the one used to create the ticket")
		}
	})
	t.Run("Should return an error if the ticket ID is nil", func(t *testing.T) {
		t.Parallel()
		_, err := agent.GetTicket(uuid.Nil)
		if err == nil {
			t.Error("Error should not be nil")
		}
		if !errors.Is(err, ErrNilTicketID) {
			t.Errorf("Error should be %v", ErrNilTicketID)
		}
	})
	t.Run("Should return an error if the ticket repository returns an error", func(t *testing.T) {
		t.Parallel()
		errorAgent, agentCreationErr := New(stubTicketRepository{
			forcedError: errors.New("an error occurred while retrieving the ticket"),
		})
		if agentCreationErr != nil {
			t.Fatal("An error occurred while creating the agent")
		}

		_, err := errorAgent.GetTicket(uuid.New())
		if err == nil {
			t.Error("Error should not be nil")
		}
		if !errors.Is(err, ErrRetrievingTicket) {
			t.Errorf("Error should be %v", ErrRetrievingTicket)
		}
	})
}

func TestBasicAgent_AnswerTicket(t *testing.T) {
	t.Parallel()
	agent, instanceError := InstanceAgent(uuid.New(), time.Now(), &fakeTicketRepository{})
	if instanceError != nil {
		t.Fatal("Error should be nil")
	}
	t.Run("Should add a response to the ticket", func(t *testing.T) {
		t.Parallel()
		ticketID := uuid.New()
		comment := "Test Comment"
		err := agent.AnswerTicket(ticketID, comment)
		if err != nil {
			t.Fatal("Error should be nil")
		}
		tkt, getTicketError := agent.GetTicket(ticketID)
		if getTicketError != nil {
			t.Fatal("Error should be nil")
		}
		if len(tkt.Responses()) != 1 {
			t.Error("Response should be added to the ticket")
		}
		if tkt.Responses()[0].Content() != comment {
			t.Error("Response should have the same comment as the one provided")
		}
	})
	t.Run("Should return an error if the comment is empty", func(t *testing.T) {
		t.Parallel()
		err := agent.AnswerTicket(uuid.New(), "")
		if err == nil {
			t.Error("Error should not be nil")
		}
		if err.Error() != "error while validating comment: response cannot be empty" {
			t.Error("Error message should be 'response cannot be empty'")
		}
	})
	t.Run("Should return an error if the ticket ID is nil", func(t *testing.T) {
		t.Parallel()
		err := agent.AnswerTicket(uuid.Nil, "Test Comment")
		if err == nil {
			t.Error("Error should not be nil")
		}
		if !errors.Is(err, ErrNilTicketID) {
			t.Errorf("Error should be %v", ErrNilTicketID)
		}
	})
	t.Run("Should return an error if the ticket repository returns an error", func(t *testing.T) {
		t.Parallel()
		errorAgent, agentCreationErr := New(stubTicketRepository{
			forcedError: errors.New("an error occurred while retrieving the ticket"),
		})
		if agentCreationErr != nil {
			t.Fatal("An error occurred while creating the agent")
		}
		err := errorAgent.AnswerTicket(uuid.New(), "Test Comment")
		if err == nil {
			t.Error("Error should not be nil")
		}
		if !errors.Is(err, ErrRetrievingTicket) {
			t.Errorf("Error should be %v", ErrRetrievingTicket)
		}
	})
}

type stubTicketRepository struct {
	forcedError error
}

func (s stubTicketRepository) UpdateTicket(tck ticket.Ticket) error {
	return nil
}

func (s stubTicketRepository) GetTicket(id uuid.UUID) (ticket.Ticket, error) {
	if s.forcedError != nil {
		return nil, s.forcedError
	}
	return ticket.MakeBasicTicket(id, time.Now(), ticket.Data{Title: "Test Title", Description: "Test Description", Status: "Open", Responses: nil})
}

type fakeTicketRepository struct {
	forcedError error
	tickets     map[uuid.UUID]ticket.Ticket
}

func (f *fakeTicketRepository) UpdateTicket(tck ticket.Ticket) error {
	if f.forcedError != nil {
		return f.forcedError
	}
	if f.tickets == nil {
		f.tickets = make(map[uuid.UUID]ticket.Ticket)
	}
	if _, ok := f.tickets[tck.ID()]; !ok {
		return errors.New("ticket not found")
	}
	f.tickets[tck.ID()] = tck
	return nil
}

func (f *fakeTicketRepository) GetTicket(id uuid.UUID) (ticket.Ticket, error) {
	if f.forcedError != nil {
		return nil, f.forcedError
	}
	if f.tickets == nil {
		f.tickets = make(map[uuid.UUID]ticket.Ticket)
	}
	if tck, ok := f.tickets[id]; ok {
		return tck, nil
	}
	newTicket, err := ticket.MakeBasicTicket(id, time.Now(), ticket.Data{Title: "Test Title", Description: "Test Description", Status: "Open", Responses: nil})
	if err != nil {
		return nil, fmt.Errorf("error while creating new ticket: %w", err)
	}
	f.tickets[id] = newTicket
	return newTicket, nil
}
