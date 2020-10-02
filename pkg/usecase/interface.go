package usecase

import (
	"github.com/kangbojk/go-react-fullstack/pkg/entity/account"
	"github.com/kangbojk/go-react-fullstack/pkg/entity/tenant"
	"github.com/kangbojk/go-react-fullstack/pkg/ID"
)

//Service provide multiple usecase to interact with account and tenant
type Service interface {
	Upgrade(a *account.Account) error
	AddUser(a *account.Account, count int) error

	FindUserWithEmail(email string) (*account.Account, error)
	GetAccount(i id.ID) (*account.Account, error)
	GetTenant(i id.ID) (*tenant.Tenant, error)
	GetTenantUsers(a *account.Account) (int, error)

	UpdateAccount(a *account.Account) error
	UpdateTenant(t *tenant.Tenant) error

	CreateAccount(a *account.Account) (id.ID, error)
	CreateTenant(t *tenant.Tenant) (id.ID, error)
	AssignTenantToAccount(a *account.Account, t *tenant.Tenant) error
}
