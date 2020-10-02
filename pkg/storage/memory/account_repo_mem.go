package in_memory

import (
	"github.com/kangbojk/go-react-fullstack/pkg/entity/account"
	"github.com/kangbojk/go-react-fullstack/pkg/ID"
)

//accountRepoMem in memory repo
type accountRepoMem struct {
	m map[id.ID]*account.Account
}

//NewAccountRepoMem create new repository
func NewAccountRepoMem() *accountRepoMem {
	var m = map[id.ID]*account.Account{}
	return &accountRepoMem{
		m: m,
	}
}

//Create an account
func (r *accountRepoMem) Create(e *account.Account) (id.ID, error) {
	r.m[e.ID] = e
	return e.ID, nil
}

//Get an account
func (r *accountRepoMem) Get(id id.ID) (*account.Account, error) {
	if r.m[id] == nil {
		return nil, account.ErrInvalid
	}
	return r.m[id], nil
}

func (r *accountRepoMem) FindUserWithEmail(email string) (*account.Account, error) {
	for _, acc := range r.m {
		if acc.Email == email {
			return acc, nil
		}
	}
	return nil, account.ErrInvalid
}

//Update an account
func (r *accountRepoMem) Update(e *account.Account) error {
	_, err := r.Get(e.ID)
	if err != nil {
		return err
	}
	r.m[e.ID] = e
	return nil
}

//List accounts
func (r *accountRepoMem) List() ([]*account.Account, error) {
	var d []*account.Account
	for _, j := range r.m {
		d = append(d, j)
	}
	return d, nil
}

//Delete an account
func (r *accountRepoMem) Delete(id id.ID) error {
	if r.m[id] == nil {
		return account.ErrInvalid
	}
	r.m[id] = nil
	return nil
}
