package agent

import (
	"errors"
	"github.com/google/uuid"
	"testing"
	"time"
)

func TestNewTicketAgentFactory(t *testing.T) {
	t.Parallel()

	t.Run("It should create a basicTicketAgentFactory with the provided repository", func(t *testing.T) {
		t.Parallel()
		var repository stubTicketRepository = stubTicketRepository{}
		var factory Factory
		var err error
		if err != nil {
			t.Fatalf("Error should be nil, got %v", err)
		}
		factory, err = NewTicketAgentFactory(repository)
		assertFactoryCreationWithProvidedRepository(t, factory, repository)
	})
	t.Run("It should return an error if the repository is nil", func(t *testing.T) {
		t.Parallel()
		_, err := NewTicketAgentFactory(nil)
		if err == nil {
			t.Fatal("Error should not be nil")
		}
		if !errors.Is(err, ErrNilRepository) {
			t.Fatalf("Error should be ErrNilRepository, got %v", err)
		}
	})
}

func TestBasicTicketAgentFactory_NewAgent(t *testing.T) {
	t.Parallel()
	var factory Factory
	var factoryCreationError error
	var repository stubTicketRepository = stubTicketRepository{}
	factory, factoryCreationError = NewTicketAgentFactory(repository)
	if factoryCreationError != nil {
		t.Fatalf("Error should be nil, got %v", factoryCreationError)
	}
	t.Run("It should create a new agent with the provided ticket.RepositoryAgentAccess", func(t *testing.T) {
		t.Parallel()
		agent, err := factory.NewAgent()
		if err != nil {
			t.Fatalf("Error should be nil, got %v", err)
		}
		if agent == nil {
			t.Fatal("Agent should not be nil")
		}
		if agent.(basicAgent).ticketRepository != repository {
			t.Fatal("Repository should be the same as the provided repository")
		}
	})
}

func TestBasicTicketAgentFactory_InstantiateAgent(t *testing.T) {
	t.Parallel()
	var factory Factory
	var factoryCreationError error
	var repository stubTicketRepository = stubTicketRepository{}
	factory, factoryCreationError = NewTicketAgentFactory(repository)
	if factoryCreationError != nil {
		t.Fatalf("Error should be nil, got %v", factoryCreationError)
	}
	t.Run("It should create a new agent with the provided ticket.RepositoryAgentAccess", func(t *testing.T) {
		t.Parallel()
		agentID := uuid.New()
		testTime := time.Now()
		agent, err := factory.InstantiateAgent(agentID, testTime)
		if err != nil {
			t.Fatalf("Error should be nil, got %v", err)
		}
		if agent == nil {
			t.Fatal("Agent should not be nil")
		}
		if agent.(basicAgent).ticketRepository != repository {
			t.Fatal("Repository should be the same as the provided repository")
		}

	})
}

func assertFactoryCreationWithProvidedRepository(t *testing.T, factory Factory, repository stubTicketRepository) {
	t.Helper()

	_, ok := factory.(basicTicketAgentFactory)
	if !ok {
		t.Fatal("Factory should be a basicTicketAgentFactory")
	}

	if factory.(basicTicketAgentFactory).ticketRepository != repository {
		t.Fatal("Repository should be the same as the provided repository")
	}
}
