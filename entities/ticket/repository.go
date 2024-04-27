package ticket

import "github.com/google/uuid"

type RepositoryClientAccess interface {
	RepositoryClientReader
	RepositoryClientWriter
}
type RepositoryClientReader interface {
	// GetTicket takes the client id to check that the client has access to the ticket, it should return an error
	// if the client does not have access to the ticket
	GetTicket(client, ticket uuid.UUID) (Ticket, error)
	GetAllClientTickets(client uuid.UUID) ([]Ticket, error)
	GetClientTicketCount(client uuid.UUID) (int, error)
}

type RepositoryClientWriter interface {
	SaveNewTicketForClient(userId uuid.UUID, ticket Ticket) error
}
