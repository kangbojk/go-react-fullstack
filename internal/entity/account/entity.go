// package account contains account entity and repository interface
package account

import (
	"errors"
	"github.com/kangbojk/gravitational_full_stack_challenge/pkg/ID"
	"time"
)

//Account defines the properties of an registered account
type Account struct {
	ID        id.ID     `json:"id"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	Actions   string    `json:"actions"`
	TenantID  id.ID     `json:"tenant_id"`
	CreatedAt time.Time `json:"created_time"`
	UpdatedAt time.Time `json:"updated_time"`
}

// Repository provides access to the account storage.
type Repository interface {
	Get(id id.ID) (*Account, error)
	List() ([]*Account, error)

	Create(e *Account) (id.ID, error)
	Update(e *Account) error
	Delete(id id.ID) error
}

// ErrInvalid is used when an account could not be found/invalid.
var ErrInvalid = errors.New("Account invalid error")
var ErrNotAuthorize = errors.New("Account not authorized")
