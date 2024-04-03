package entities

import (
	"github.com/rogelioConsejo/golibs/helpers"
	"testing"
	"time"
)

func TestClient(t *testing.T) {
	t.Parallel()
	client := makeBasicClient(t)
	title := helpers.MakeRandomString(10)
	description := helpers.MakeRandomString(100)

	t.Run("A basic client can be created and has a creation date", func(t *testing.T) {
		t.Parallel()
		checkClientCreationTime(t, client)
	})
	t.Run("A Client can create tickets", func(t *testing.T) {
		err := client.CreateTicket(title, description)
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}
		if client.TicketCount() != 1 {
			t.Errorf("Expected ticket count to be 1, got %v", client.TicketCount())
		}
		_ = client.CreateTicket(title+" for another issue", description+" for another issue")
		if client.TicketCount() != 2 {
			t.Errorf("Expected ticket count to be 2, got %v", client.TicketCount())
		}
	})
	t.Run("A client can return all its tickets", func(t *testing.T) {
		var clientTickets []Ticket = client.GetTickets()
		if clientTickets[0].Title() != title {
			t.Errorf("Expected title to be %v, got %v", title, clientTickets[0].Title())
		}
		if clientTickets[0].Description() != description {
			t.Errorf("Expected description to be %v, got %v", description, clientTickets[0].Description())
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
	client := NewBasicTicketClient()
	if client == nil {
		t.Errorf("Expected a client, got nil")
	}
	return client

}
