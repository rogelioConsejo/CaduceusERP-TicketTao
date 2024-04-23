package client

import (
	"github.com/rogelioConsejo/golibs/helpers"
	"testing"
	"ticketTao/entities/ticket"
	"time"
)

// TestClient is a test function to test the functionality of the TicketClient. It runs multiple sub-tests to cover different scenarios.
// The test function creates a basic client using the makeBasicClient function and performs the following tests:
//   - A basic client can be created and has a creation date
//   - A client can create tickets and the ticket count is validated
//   - A client can return all its tickets and validates the title and description of the first ticket
//   - A client can return a specific ticket by index and validates the title and description of the ticket
//
// The test function requires the makeBasicClient function and the checkClientCreationTime function.
//
// For running the sub-tests in parallel, the t.Parallel() function is used.
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
		var clientTickets []ticket.Ticket = client.GetTickets()
		if clientTickets[0].Title() != title {
			t.Errorf("Expected title to be %v, got %v", title, clientTickets[0].Title())
		}
		if clientTickets[0].Description() != description {
			t.Errorf("Expected description to be %v, got %v", description, clientTickets[0].Description())
		}
	})
	t.Run("A client can return a specific ticket by index", func(t *testing.T) {
		if client.GetTicket(1).Title() != title+" for another issue" {
			t.Errorf("Expected title to be %v, got %v", title+" for another issue", client.GetTicket(1).Title())
		}
		if client.GetTicket(1).Description() != description+" for another issue" {
			t.Errorf("Expected description to be %v, got %v", description+" for another issue", client.GetTicket(1).Description())
		}

	})
}

// checkClientCreationTime checks if the creation date of a client is in the past
func checkClientCreationTime(t *testing.T, client TicketClient) {
	t.Helper()
	time.Sleep(11 * time.Millisecond)
	if client.CreatedAt().After(time.Now().Add(-10 * time.Millisecond)) {
		t.Errorf("Expected creation date to be in the past, got %v", client.CreatedAt())
	}
}

// makeBasicClient is a helper function that creates and returns a `TicketClient` using the `NewBasicTicketClient` function.
func makeBasicClient(t *testing.T) TicketClient {
	t.Helper()
	client := NewBasicTicketClient()
	if client == nil {
		t.Errorf("Expected a client, got nil")
	}
	return client
}
