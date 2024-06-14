package agent

import (
	"github.com/google/uuid"
	"testing"
)

func TestTicketCloserAgent_CloseTicket(t *testing.T) {
	t.Parallel()
	t.Run("It should close the ticket", func(t *testing.T) {
		repository := stubTicketRepository{}
		factory, factoryCreationError := NewTicketAgentFactory(repository)
		if factoryCreationError != nil {
			t.Fatalf("Error should be nil, got %v", factoryCreationError)
		}
		agent, err := factory.NewAgent()
		if err != nil {
			t.Fatalf("Error should be nil, got %v", err)
		}
		closerAgent, err := factory.InstantiateTicketCloserAgent(agent.ID(), agent.CreatedAt())
		if err != nil {
			t.Fatalf("Error should be nil, got %v", err)
		}
		ticketID := uuid.New()
		err = closerAgent.CloseTicket(ticketID)
		t.Parallel()
	})
}
