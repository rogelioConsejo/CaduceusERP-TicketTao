package entities

import "time"

type CreatedEntity interface {
	CreatedAt() time.Time
}
