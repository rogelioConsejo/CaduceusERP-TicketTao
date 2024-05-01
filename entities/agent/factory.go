package agent

import (
	"errors"
	"fmt"
	"github.com/google/uuid"
	"ticketTao/entities/ticket"
	"time"
)

func NewTicketAgentFactory(repository ticket.RepositoryAgentAccess) (Factory, error) {
	if repository == nil {
		return nil, ErrNilRepository

	}
	return basicTicketAgentFactory{
		ticketRepository: repository,
	}, nil
}

type Factory interface {
	NewAgent() (Agent, error)
	InstantiateAgent(agent uuid.UUID, createdAt time.Time) (Agent, error)
}

type basicTicketAgentFactory struct {
	ticketRepository ticket.RepositoryAgentAccess
}

func (b basicTicketAgentFactory) InstantiateAgent(agent uuid.UUID, createdAt time.Time) (Agent, error) {
	newAgent, err := InstanceAgent(agent, createdAt, b.ticketRepository)
	if err != nil {
		return nil, fmt.Errorf("%w: %w", ErrInstantiatingAgent, err)

	}
	return newAgent, nil
}

func (b basicTicketAgentFactory) NewAgent() (Agent, error) {
	newAgent, err := New(b.ticketRepository)
	if err != nil {
		return nil, fmt.Errorf("%w: %w", ErrCreatingAgent, err)
	}
	return newAgent, nil
}

var ErrNilRepository = errors.New("repository can't be nil")
var ErrCreatingAgent = errors.New("error creating agent")
var ErrInstantiatingAgent = errors.New("error instantiating agent")
