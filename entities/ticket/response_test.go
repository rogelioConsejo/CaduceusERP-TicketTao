package ticket

import (
	"github.com/google/uuid"
	"testing"
	"time"
)

func TestNewResponse(t *testing.T) {
	// TestNewResponse tests the creation of a new basicResponse object.
	var id uuid.UUID = uuid.New()
	const content string = "test content"
	var response Response = NewResponse(id, content)

	assertResponseProperties(t, id, content, response)

	t.Run("A response has a timestamp", func(t *testing.T) {
		t.Parallel()
		assertResponseTimestamp(t, response)
	})
}

func TestMakeResponse(t *testing.T) {
	// TestMakeResponse tests the creation of a new basicResponse object with a given timestamp.
	var id uuid.UUID = uuid.New()
	const content string = "test content"
	var timeStamp time.Time = time.Now().Add(-10 * time.Second)
	var response Response = MakeResponse(id, content, timeStamp)

	assertResponseProperties(t, id, content, response)

	t.Run("A response has a timestamp", func(t *testing.T) {
		t.Parallel()
		if !response.TimeStamp().Equal(timeStamp) {
			t.Errorf("Expected timestamp to be %v, got %v", timeStamp, response.TimeStamp())
		}
	})
}

func assertResponseTimestamp(t *testing.T, response Response) {
	t.Helper()
	timeStamp := response.TimeStamp()
	if timeStamp.IsZero() {
		t.Error("Expected timestamp to be set")
	}
	if timeStamp.After(time.Now()) {
		t.Error("Expected timestamp to be in the past")
	}
	time.Sleep(1 * time.Millisecond)
	if !response.TimeStamp().Equal(timeStamp) {
		t.Errorf("Expected timestamp to be %v, got %v", timeStamp, response.TimeStamp())
	}
}

func assertResponseProperties(t *testing.T, id uuid.UUID, content string, response Response) {
	t.Helper()
	// assert that the ID method returns a UUID
	var _ uuid.UUID = response.UserId()
	// assert that the returned user ID is the same as the one passed to the constructor
	if response.UserId() != id {
		t.Errorf("Expected user ID to be %v, got %v", id, response.UserId())
	}
	// assert that the content of the response is the same as the one passed to the constructor
	if response.Content() != content {
		t.Errorf("Expected content to be %v, got %v", content, response.Content())
	}
}
