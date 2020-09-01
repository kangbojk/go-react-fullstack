package tenant

import (
	"time"

	"gravitational_full_stack_challenge/pkg/ID"
)

func NewFixtureTenant() *Tenant {
	return &Tenant{
		ID:        id.NewID(),
		Users:     0,
		Capacity:  100,
		CreatedAt: time.Now(),
	}
}
