package tenant

import (
	"time"

	"github.com/kangbojk/go-react-fullstack/pkg/ID"
)

func NewFixtureTenant() *Tenant {
	return &Tenant{
		ID:        id.NewID(),
		Users:     0,
		Capacity:  100,
		CreatedAt: time.Now(),
	}
}
