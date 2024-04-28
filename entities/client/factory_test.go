package client

import (
	"github.com/google/uuid"
	"testing"
	"time"
)

func TestFactory(t *testing.T) {
	t.Parallel()

	repo := makeFakeTicketRepository()
	var clientFactory Factory = NewClientFactory(repo)

	t.Run("A factory can create new clients", func(t *testing.T) {
		client := clientFactory.NewBasicTicketClient()
		assertClientIsNotNil(t, client)
		assertNewClientValues(t, client)
	})
	t.Run("A factory can instantiate clients", func(t *testing.T) {
		id := uuid.New()
		creationTime := time.Now()
		client := clientFactory.InstantiateBasicTicketClient(id, creationTime)
		assertClientIsNotNil(t, client)
		assertClientValues(t, client, id, creationTime)
	})
}

func assertClientValues(t *testing.T, client TicketClient, id uuid.UUID, creationTime time.Time) {
	t.Helper()
	if client.ID() != id {
		t.Errorf("Expected the client to have id %s, got %s", id.String(), client.ID().String())
	}
	if client.CreatedAt() != creationTime {
		t.Errorf("Expected the client to have creation time %v, got %v", creationTime, client.CreatedAt())
	}
}

func assertNewClientValues(t *testing.T, client TicketClient) {
	t.Helper()
	if client.ID() == uuid.Nil {
		t.Error("Expected the client to have an id")
	}
	if client.CreatedAt().IsZero() {
		t.Error("Expected the client to have a creation time")
	}
	if client.CreatedAt().After(time.Now()) {
		t.Error("Expected the creation time to be in the past")
	}
	count, err := client.TicketCount()
	if err != nil {
		t.Errorf("Expected no error retrieving ticket count, got %v", err)
	}
	if count != 0 {
		t.Error("Expected the new client to have no tickets")
	}
	tickets, err := client.GetTickets()
	if err != nil {
		t.Errorf("Expected no error retrieving tickets, got %v", err)
	}
	if len(tickets) != 0 {
		t.Error("Expected the new client to have no tickets")
	}
}

func assertClientIsNotNil(t *testing.T, client TicketClient) {
	t.Helper()
	if client == nil {
		t.Fatal("Expected a client to be returned")
	}
}
