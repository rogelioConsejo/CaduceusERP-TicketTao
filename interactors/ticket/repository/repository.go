// Package repository contains an implementation of the ticket repository interface
package repository

import (
	"github.com/google/uuid"
	"ticketTao/entities/ticket"
)

// TicketPersistence is an interface that defines the methods that a ticket persistence driver should implement
type TicketPersistence interface {
	SaveNewTicketForClient(client uuid.UUID, tck ticket.Ticket) error
	GetTicketOwner(ticket uuid.UUID) (client uuid.UUID, err error)
	GetTicket(id uuid.UUID) (ticket.Ticket, error)
}
