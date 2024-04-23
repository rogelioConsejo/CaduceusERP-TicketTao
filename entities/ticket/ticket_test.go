package ticket

import (
	"github.com/google/uuid"
	"github.com/rogelioConsejo/golibs/helpers"
	"testing"
	"time"
)

func TestMakeEmptyBasicTicket(t *testing.T) {
	t.Parallel()

	t.Run("A basic ticket can be created with an ID, title, and description", func(t *testing.T) {
		t.Parallel()
		id := uuid.New()
		title := helpers.MakeRandomString(10)
		description := helpers.MakeRandomString(100)
		ticket := MakeEmptyBasicTicket(id, title, description)
		assertEqual(t, "title", ticket.Title(), title)
		assertEqual(t, "description", ticket.Description(), description)
		assertEqual(t, "id", ticket.ID(), id)
		checkTicketCreationTime(t, ticket)
	})
}

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
		assertEqual(t, "title", ticket.Title(), title)
		assertEqual(t, "description", ticket.Description(), description)
		checkTicketCreationTime(t, ticket)
	})

	t.Run("A ticket has an ID", func(t *testing.T) {
		t.Parallel()
		id := uuid.New()
		title := "A title"
		description := "A description"
		ticket := MakeEmptyBasicTicket(id, title, description)
		assertEqual(t, "id", ticket.ID(), id)
	})

	t.Run("A ticket is created with an 'Open' status", func(t *testing.T) {
		t.Parallel()
		ticket := makeBasicTicket(t)
		assertEqual(t, "status", ticket.Status(), Open)
	})
}

func TestBasicTicket_AddResponse(t *testing.T) {
	t.Parallel()
	ticket, response := setupTicketAndResponse(t)
	addResponseAndCheck(t, ticket, response)

	t.Run("A ticket status changes to 'in progress' when a response is added", func(t *testing.T) {
		t.Parallel()
		checkInProgressStatus(t, ticket)
	})
}

func TestBasicTicket_Close(t *testing.T) {
	t.Parallel()
	ticket := makeBasicTicket(t)
	ticket.Close()
	var status = ticket.Status()
	if status != Closed {
		t.Errorf("Expected status to be 'Closed', got %v", status)
	}
}

func assertEqual(t *testing.T, field string, got, want interface{}) {
	t.Helper()
	if got != want {
		t.Errorf("Expected %s to be %v, got %v", field, want, got)
	}
}

func setupTicketAndResponse(t *testing.T) (Ticket, Response) {
	t.Helper()
	ticket := makeBasicTicket(t)
	userId := uuid.New()
	responseText := helpers.MakeRandomString(10)
	var response = NewResponse(userId, responseText)
	return ticket, response
}

func addResponseAndCheck(t *testing.T, ticket Ticket, response Response) {
	t.Helper()
	ticket.AddResponse(response)
	if len(ticket.Responses()) != 1 {
		t.Fatalf("Expected 1 response, got %v", len(ticket.Responses()))
	}
	if ticket.Responses()[0] != response {
		t.Errorf("Expected response to be %v, got %v", response, ticket.Responses()[0])
	}
}

func checkInProgressStatus(t *testing.T, ticket Ticket) {
	t.Helper()
	if ticket.Status() != InProgress {
		t.Errorf("Expected status to be 'in progress', got %v", ticket.Status())
	}
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
