package support

import (
	"github.com/google/uuid"
	"testing"
	"ticketTao/entities/ticket"
	"time"
)

func TestNewSupportAgent(t *testing.T) {
	var agent Agent = NewAgent()

	// assert that the ID UUID is not empty
	if agent.ID().String() == "" {
		t.Error("ID should not be empty")
	}
	// assert that the ID method returns a UUID
	var _ uuid.UUID = agent.ID()

	t.Run("An agent should have a creation time", func(t *testing.T) {
		if agent.CreatedAt().IsZero() {
			t.Error("Creation time should not be zero")
		}
		if agent.CreatedAt().After(time.Now()) {
			t.Error("Creation time should be before now")
		}
	})
}

func TestInstanceSupportAgent(t *testing.T) {
	var uuid1 uuid.UUID = uuid.New()
	var agent Agent = InstanceSupportAgent(uuid1)
	if agent.ID() != uuid1 {
		t.Error("ID should be equal to the one used to create the agent")
	}
}

func TestBasicSupportAgent_GetTicket(t *testing.T) {
	agent := basicAgent{
		id:               uuid.New(),
		creationTime:     time.Now(),
		ticketRepository: stubTicketRepository{},
	}

	testUUID := uuid.New()
	tkt := agent.GetTicket(testUUID)
	if tkt.ID() != testUUID {
		t.Error("Ticket ID should be equal to the one used to create the ticket")
	}
}

type stubTicketRepository struct {
}

func (s stubTicketRepository) Get(id uuid.UUID) ticket.Ticket {
	return ticket.MakeEmptyBasicTicket(id, "Test Title", "Test Description")
}
