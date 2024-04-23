package ticket

import (
	"github.com/google/uuid"
	"time"
)

type Response interface {
	UserId() uuid.UUID
	Content() string
	TimeStamp() time.Time
}

// NewResponse is a function that creates a new basicResponse object with the given user ID and content.
// The function takes a user ID of type uuid.UUID and a content string, and returns a basicResponse object.
func NewResponse(id uuid.UUID, s string) Response {
	return basicResponse{
		userId:    id,
		content:   s,
		timeStamp: time.Now(),
	}
}

func MakeResponse(id uuid.UUID, s string, t time.Time) Response {
	return basicResponse{
		userId:    id,
		content:   s,
		timeStamp: t,
	}
}

// basicResponse represents a response object with a user ID and content.
type basicResponse struct {
	timeStamp time.Time
	userId    uuid.UUID
	content   string
}

// UserId returns the user ID of the user who created the response.
func (r basicResponse) UserId() uuid.UUID {
	return r.userId
}

func (r basicResponse) Content() string {
	return r.content
}

// TimeStamp returns the time when the response was created.
func (r basicResponse) TimeStamp() time.Time {
	return r.timeStamp
}
