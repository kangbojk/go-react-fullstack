package account

import (
	// "strings"
	"time"

	"github.com/kangbojk/go-react-fullstack/pkg/ID"
)

type manager struct {
	repo Repository
}

//NewManager create new manager
func NewManager(r Repository) *manager {
	return &manager{
		repo: r,
	}
}

//Create an account
func (s *manager) Create(e *Account) (id.ID, error) {
	e.ID = id.NewID()
	e.CreatedAt = time.Now()
	return s.repo.Create(e)
}

//Get an account
func (s *manager) Get(id id.ID) (*Account, error) {
	return s.repo.Get(id)
}

//List Accounts
func (s *manager) List() ([]*Account, error) {
	return s.repo.List()
}

//Delete an account
func (s *manager) Delete(id id.ID) error {
	_, err := s.Get(id)
	if err != nil {
		return err
	}
	return s.repo.Delete(id)
}

//Update an account
func (s *manager) Update(e *Account) error {
	return s.repo.Update(e)
}
