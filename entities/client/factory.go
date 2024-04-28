package client

import (
	"github.com/google/uuid"
	"ticketTao/entities/ticket"
	"time"
)

func NewClientFactory(tr ticket.RepositoryClientAccess) Factory {
	return basicTicketClientFactory{
		ticketRepository: tr,
	}
}

type Factory interface {
	NewBasicTicketClient() TicketClient
	InstantiateBasicTicketClient(client uuid.UUID, time time.Time) TicketClient
}

type basicTicketClientFactory struct {
	ticketRepository ticket.RepositoryClientAccess
}

func (b basicTicketClientFactory) InstantiateBasicTicketClient(client uuid.UUID, time time.Time) TicketClient {
	return InstantiateBasicTicketClient(client, time, b.ticketRepository)
}

func (b basicTicketClientFactory) NewBasicTicketClient() TicketClient {
	return NewBasicTicketClient(b.ticketRepository)
}
