// package usecase implements service interface
package usecase

import (
	"time"

	"github.com/kangbojk/go-react-fullstack/pkg/entity/account"
	"github.com/kangbojk/go-react-fullstack/pkg/entity/tenant"
	"github.com/kangbojk/go-react-fullstack/pkg/ID"
)

type service struct {
	aRepo account.Repository
	tRepo tenant.Repository
}

//NewService create new service
func NewService(accountRepo account.Repository, tenantRepo tenant.Repository) *service {
	return &service{
		aRepo: accountRepo,
		tRepo: tenantRepo,
	}
}

func (s *service) Upgrade(userAccount *account.Account) error {
	userAccount, err := s.aRepo.Get(userAccount.ID)
	if err != nil {
		return err
	}

	if userAccount.Actions != "RW" {
		return account.ErrNotAuthorize
	}

	userTenant, err := s.tRepo.Get(userAccount.TenantID)
	if err != nil {
		return err
	}

	if userTenant.Capacity >= 1000 {
		return tenant.ErrAlreadyUpgraded
	}

	userTenant.Capacity = 1000
	err = s.tRepo.Update(userTenant)
	if err != nil {
		return err
	}

	return nil
}

func (s *service) AddUser(a *account.Account, count int) error {
	userAccount, err := s.aRepo.Get(a.ID)
	if err != nil {
		return err
	}

	if userAccount.Actions != "RW" {
		return account.ErrNotAuthorize
	}

	userTenant, err := s.tRepo.Get(userAccount.TenantID)
	if err != nil {
		return err
	}

	// var full error
	if userTenant.Users+count <= userTenant.Capacity {
		userTenant.Users += count
		// } else if userTenant.Users+count == userTenant.Capacity {
		// 	userTenant.Users += count
		// 	full = tenant.ErrFull
	} else {
		return tenant.ErrFull
	}

	err = s.tRepo.Update(userTenant)
	if err != nil {
		return err
	}

	// if full != nil {
	// 	return full
	// }
	return nil
}

// GetTenantUsers Retrieve user count in tenant
func (s *service) GetTenantUsers(userAccount *account.Account) (int, error) {
	userAccount, err := s.aRepo.Get(userAccount.ID)
	if err != nil {
		return -1, err
	}

	userTenant, err := s.tRepo.Get(userAccount.TenantID)
	if err != nil {
		return -1, err
	}

	return userTenant.Users, nil
}

// CreateAccount Create an account
func (s *service) CreateAccount(a *account.Account) (id.ID, error) {
	a.ID = id.NewID()
	a.CreatedAt = time.Now()
	return s.aRepo.Create(a)
}

// CreateTenant Create a tenant
func (s *service) CreateTenant(t *tenant.Tenant) (id.ID, error) {
	t.ID = id.NewID()
	t.CreatedAt = time.Now()
	return s.tRepo.Create(t)
}

func (s *service) GetAccount(i id.ID) (*account.Account, error) {
	return s.aRepo.Get(i)
}

func (s *service) GetTenant(i id.ID) (*tenant.Tenant, error) {
	return s.tRepo.Get(i)
}

func (s *service) AssignTenantToAccount(a *account.Account, t *tenant.Tenant) error {
	a.TenantID = t.ID
	return s.aRepo.Update(a)
}

func (s *service) FindUserWithEmail(email string) (*account.Account, error) {
	return s.aRepo.FindUserWithEmail(email)
}

func (s *service) UpdateAccount(a *account.Account) error {
	return s.aRepo.Update(a)
}

func (s *service) UpdateTenant(t *tenant.Tenant) error {
	return s.tRepo.Update(t)
}
