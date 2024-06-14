# TicketTao for CaduceusERP
## Naturally flowing ticketing system.
"TicketTao for CaduceusERP" is more than just a ticketing system; it's a philosophy towards problem-solving and task 
management, encouraging a more Zen-like approach to addressing issues that arise in the business environment, fostering 
a calm, organized, and efficient workplace.

### Flow State in Work
Inspired by the Taoist concept of effortless action (Wu Wei), creating a ticketing system that encourages a flow state 
among users is paramount. Flow state, a concept identified by psychologist Mihaly Csikszentmihalyi, refers to a state of 
immersion and focused energy on tasks. "TicketTao" aims to minimize disruptions and cognitive load, allowing users 
to maintain focus and achieve a higher productivity level.

### Scalable Simplicity
Scalability doesn’t just pertain to handling growing amounts of work or data; it also involves maintaining simplicity 
in the face of complexity. As businesses grow, their processes and systems tend to become more complicated. A system 
like "TicketTao" is designed to scale not only in capacity but also in maintaining or even simplifying the user 
experience as demands increase.

### Intuitive Design
Drawing from the Taoist principle of following the natural path, "TicketTao" will prioritize intuitive design, ensuring 
that users can navigate and utilize the system with minimal training. The goal is to align the system’s functionality 
with the users’ innate behaviors and expectations, reducing resistance and enhancing adoption rates.

> Simple tickets flow,  
> CaduceusERP's glow,  
> Tasks resolve, stress low.  

## Use cases:
### v0.0.1a
- A ticket is created with a title and description.
- A ticket is created with an initial "open" status
- A basic client can be created and has a creation date.
- A client can create tickets.
- A client can return all its tickets.
- A client can return a specific ticket by index.
- A support agent can be created and has a creation date and a unique UUID.
- A support agent can be instantiated by it UUID.
- A support agent has a GetTicket method to retrieve tickets.

#### Development
##### Must have
- [x] A ticket changes to "in progress" status when a support agent answers it
- [x] A ticket can be closed
- [x] A ticket has a conversation log with timestamp, user id and message
- [x] A ticket has a creation date
- [x] A ticket has a unique identifier
- [ ] A ticket repository can save a ticket for a client
- [ ] A ticket repository can return a ticket by its id
- [ ] A ticket repository can return all of a client's tickets ordered by creation date
- [ ] A ticket repository can return all tickets ordered by creation date (fifo)
- [ ] A ticket repository can return only non-closed tickets
- [ ] A ticket repository can close tickets
- [x] A client uses a ticket repository for ticket persistence
- [x] A client can add an answer or comment to a ticket
- [x] A client can close a ticket
- [x] A client can retrieve all their tickets through the ticket repository
- [x] A client can get a specific ticket by its id
- [x] A support agent uses a ticket repository for ticket persistence
- [x] A support agent can add an answer to a ticket
- [x] A support agent can be allowed to close tickets though the ticket repository*
- [x] A support agent can see the ticket details

*We want to have the option of allowing support agents to close tickets or not. 

#### Next Steps 
- A client can edit their tickets
- A support agent can be enabled to create tickets
- Input validation
- Error handling
- Authentication and authorization
- User sessions
- Audit logs (limit lifespan or visibility, or get consent)
- Notifications and alerts
- Concurrency incidents management/resolution
- High priority tickets
- Low priority tickets
- Ticket visibility levels
- A history of ticket changes per ticket