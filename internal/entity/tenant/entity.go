package tenant

import (
	"time"
	"errors"
	"github.com/kangbojk/go-react-fullstack/pkg/ID"
)

//Tenant data
type Tenant struct {
	ID        id.ID     `json:"id"`
	Users     int       `json:"users"`
	Capacity  int       `json:"capacity"`
	CreatedAt time.Time `json:"created_time"`
	UpdatedAt time.Time `json:"updated_time"`
}

//Repository interface
type Repository interface {
	Get(id id.ID) (*Tenant, error)
	List() ([]*Tenant, error)

	Create(e *Tenant) (id.ID, error)
	Update(e *Tenant) error
	Delete(id id.ID) error
}

var ErrAlreadyUpgraded = errors.New("Plan already upgraded")
var ErrFull = errors.New("Tenant full, please upgrade")
var ErrNotFound = errors.New("No tenant")