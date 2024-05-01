package agent

import (
	"errors"
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

type stubTicketRepository struct {
	forcedError error
}

func (s stubTicketRepository) GetTicket(id uuid.UUID) (ticket.Ticket, error) {
	if s.forcedError != nil {
		return nil, s.forcedError
	}
	return ticket.MakeEmptyBasicTicket(id, "Test Title", "Test Description"), nil
}
