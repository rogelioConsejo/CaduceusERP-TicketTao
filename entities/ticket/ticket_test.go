package ticket

import (
	"errors"
	"github.com/google/uuid"
	"github.com/rogelioConsejo/golibs/helpers"
	"testing"
	"ticketTao/entities"
	"time"
)

func TestMakeEmptyBasicTicket(t *testing.T) {
	t.Parallel()

	t.Run("A basic ticket can be created with an ID, title, and description", func(t *testing.T) {
		t.Parallel()
		title := helpers.MakeRandomString(10)
		description := helpers.MakeRandomString(100)
		ticket, _ := NewBasicTicket(title, description)
		assertEqual(t, "title", ticket.Title(), title)
		assertEqual(t, "description", ticket.Description(), description)
		assertIdNotNil(t, ticket)
		checkTicketCreationTime(t, ticket)
	})
	t.Run("A basic ticket cannot be created with an empty title", func(t *testing.T) {
		t.Parallel()
		title := ""
		description := helpers.MakeRandomString(100)
		_, err := NewBasicTicket(title, description)
		if err == nil {
			t.Error("Expected an error, got nil")
		}
		assertErrors(t, err, NewBasicTicketError, ErrEmptyTitle)
	})
}

func TestMakeBasicTicket(t *testing.T) {
	t.Parallel()

	id := uuid.New()
	creationTime := time.Now()
	title := helpers.MakeRandomString(10)
	description := helpers.MakeRandomString(100)
	status := Status(helpers.MakeRandomString(10))
	responses := makeFakeResponses(8)

	t.Run("A basic ticket can be represented with and id, a creation time and ticket data", func(t *testing.T) {
		t.Parallel()

		var data Data
		data.Title = title
		data.Description = description
		data.Status = status
		data.Responses = responses

		tck, err := MakeBasicTicket(id, creationTime, data)
		if err != nil {
			t.Fatalf("Error creating basic ticket: %v", err)
		}

		assertEqual(t, "id", tck.ID(), id)
		assertEqual(t, "creation time", tck.CreatedAt(), creationTime)
		assertEqual(t, "title", tck.Title(), title)
		assertEqual(t, "description", tck.Description(), description)
		assertEqual(t, "status", tck.Status(), status)
		assertEqualArrays(t, "responses", tck.Responses(), responses)
	})

	t.Run("A basic ticket cannot be created with an empty title", func(t *testing.T) {
		t.Parallel()

		var data Data
		data.Title = ""
		data.Description = description
		data.Status = status
		data.Responses = responses

		_, err := MakeBasicTicket(id, creationTime, data)

		assertErrors(t, err, NewBasicTicketError, ErrEmptyTitle)
	})

	t.Run("A basic ticket cannot be created with an empty status", func(t *testing.T) {
		t.Parallel()

		var data Data
		data.Title = title
		data.Description = description
		data.Status = ""
		data.Responses = responses

		_, err := MakeBasicTicket(id, creationTime, data)

		assertErrors(t, err, NewBasicTicketError, ErrEmptyStatus)
	})

	t.Run("A basic ticket can be created with an empty responses array", func(t *testing.T) {
		t.Parallel()

		var data Data
		data.Title = title
		data.Description = description
		data.Status = status
		data.Responses = []Response{}

		tck, err := MakeBasicTicket(id, creationTime, data)
		if err != nil {
			t.Fatalf("Error creating basic ticket: %v", err)
		}

		assertEqualArrays(t, "responses", tck.Responses(), []Response{})
	})

	t.Run("A basic ticket cannot be created with a nil id", func(t *testing.T) {
		t.Parallel()

		var data Data
		data.Title = title
		data.Description = description
		data.Status = status
		data.Responses = responses

		_, err := MakeBasicTicket(uuid.Nil, creationTime, data)

		assertErrors(t, err, NewBasicTicketError, entities.ErrNilID)
	})

	t.Run("A basic ticket cannot be created with a nil creation date", func(t *testing.T) {
		t.Parallel()

		var data Data
		data.Title = title
		data.Description = description
		data.Status = status
		data.Responses = responses

		_, err := MakeBasicTicket(id, time.Time{}, data)

		assertErrors(t, err, NewBasicTicketError, entities.ErrNilCreationTime)
	})

	t.Run("A basic ticket cannot be created with creation time in the future", func(t *testing.T) {
		t.Parallel()

		var data Data
		data.Title = title
		data.Description = description
		data.Status = status
		data.Responses = responses

		_, err := MakeBasicTicket(id, time.Now().Add(10*time.Second), data)

		assertErrors(t, err, NewBasicTicketError, entities.ErrFutureCreationTime)
	})
}

func assertEqualArrays[T comparable](t *testing.T, field string, got, want []T) {
	t.Helper()
	if len(got) != len(want) {
		t.Errorf("Expected %s to have length %v, got %v", field, len(want), len(got))
	}
	for i, v := range got {
		if v != want[i] {
			t.Errorf("Expected %s[%v] to be %v, got %v", field, i, want[i], v)
		}
	}

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
		ticket, _ := NewBasicTicket(title, description)
		assertEqual(t, "title", ticket.Title(), title)
		assertEqual(t, "description", ticket.Description(), description)
		checkTicketCreationTime(t, ticket)
	})

	t.Run("A ticket has an ID", func(t *testing.T) {
		t.Parallel()
		title := "A title"
		description := "A description"
		ticket, _ := NewBasicTicket(title, description)
		assertIdNotNil(t, ticket)
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

func assertIdNotNil(t *testing.T, ticket Ticket) {
	t.Helper()
	if ticket.ID() == uuid.Nil {
		t.Errorf("Expected ID to not be nil, got %v", ticket.ID())
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
	ticket, _ := NewBasicTicket(title, description)
	if ticket == nil {
		t.Errorf("Expected a ticket, got nil")
	}
	return ticket
}

func makeFakeResponses(i int) []Response {
	var responses []Response
	for j := 0; j < i; j++ {
		clientUserId := uuid.New()
		agentUserId := uuid.New()
		content := helpers.MakeRandomString(10)

		if j%2 == 0 {
			responses = append(responses, NewResponse(clientUserId, content))
		} else {
			responses = append(responses, NewResponse(agentUserId, content))
		}

	}
	return responses

}

func assertErrors(t *testing.T, err error, expected ...error) {
	t.Helper()
	if err == nil {
		t.Error("Expected an error, got nil")
	}
	for _, e := range expected {
		if !errors.Is(err, e) {
			t.Errorf("Expected error to be %v, got %v", e, err)
		}
	}
}
