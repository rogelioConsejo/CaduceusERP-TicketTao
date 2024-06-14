package agent

import (
	"errors"
	"fmt"
	"github.com/google/uuid"
	"ticketTao/entities/ticket"
)

type TicketCloserAgent interface {
	Agent
	CloseTicket(ticket uuid.UUID) error
}

func newTicketCloserAgent(agent Agent, repo ticket.RepositoryAgentAccess) TicketCloserAgent {
	return ticketCloserAgent{
		agent,
		repo,
	}
}

type ticketCloserAgent struct {
	Agent
	repo ticket.RepositoryAgentAccess
}

func (t ticketCloserAgent) CloseTicket(id uuid.UUID) error {
	tck, err := t.repo.GetTicket(id)
	if err != nil {
		return fmt.Errorf("%w: %w", TicketRetrievalError, err)
	}
	tck.Close()
	err = t.repo.UpdateTicket(tck)
	return nil
}

var TicketRetrievalError = errors.New("error while retrieving ticket")
