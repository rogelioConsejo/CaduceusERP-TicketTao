package ticket

import "github.com/google/uuid"

// RepositoryClientAccess is an interface that defines the methods that a ticket repository should implement
// to be used by a client.
type RepositoryClientAccess interface {
	RepositoryClientReader
	RepositoryClientWriter
}

// RepositoryAgentAccess is an interface that defines the methods that a ticket repository should implement
// to be used by an agent.
type RepositoryAgentAccess interface {
	// GetTicket can retrieve a ticket from the repository based on the provided UUID. It should have access to all tickets.
	GetTicket(ticket uuid.UUID) (Ticket, error)
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
