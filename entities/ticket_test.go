package entities

import (
	"github.com/rogelioConsejo/golibs/helpers"
	"testing"
	"time"
)

func TestTicket(t *testing.T) {
	t.Parallel()
	t.Run("A basic ticket can be created and has a creation date", func(t *testing.T) {
		t.Parallel()
		ticket := makeBasicTicket(t)

		checkTicketCreationTime(t, ticket)
	})
	t.Run("A ticket has a title and a description", func(t *testing.T) {
		t.Parallel()
		title := helpers.MakeRandomString(10)
		description := helpers.MakeRandomString(100)
		ticket := NewBasicTicket(title, description)
		if ticket.Title() != title {
			t.Errorf("Expected title to be %v, got %v", title, ticket.Title())
		}
		if ticket.Description() != description {
			t.Errorf("Expected description to be %v, got %v", description, ticket.Description())
		}
		checkTicketCreationTime(t, ticket)

	})
}

func checkTicketCreationTime(t *testing.T, ticket Ticket) {
	t.Helper()
	time.Sleep(11 * time.Millisecond)
	if ticket.CreatedAt().After(time.Now().Add(-10 * time.Millisecond)) {
		t.Errorf("Expected creation date to be in the past, got %v", ticket.CreatedAt())
	}
}

func makeBasicTicket(t *testing.T) Ticket {
	t.Helper()
	const title = "A title"
	const description = "A description"
	ticket := NewBasicTicket(title, description)
	if ticket == nil {
		t.Errorf("Expected a ticket, got nil")
	}
	return ticket
}
