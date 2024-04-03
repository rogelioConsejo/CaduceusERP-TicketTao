package entities

import "time"

func NewBasicTicket(title, description string) Ticket {
	return basicTicket{
		creationTime: time.Now(),
		title:        title,
		description:  description,
	}
}

type Ticket interface {
	CreatedEntity
	Title() string
	Description() string
}

type basicTicket struct {
	creationTime time.Time
	title        string
	description  string
}

func (b basicTicket) Title() string {
	return b.title
}

func (b basicTicket) Description() string {
	return b.description
}

func (b basicTicket) CreatedAt() time.Time {
	return b.creationTime
}
