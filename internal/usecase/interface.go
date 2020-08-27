package usecase

import (
	"github.com/kangbojk/gravitational_full_stack_challenge/domain/entity/account"
	"github.com/kangbojk/gravitational_full_stack_challenge/domain/entity/tenant"
)

//Service provide multiple usecase to interact with account and tenant
type Service interface {
	Upgrade(a *account.Account) error
	AddUser(a *account.Account, t *tenant.Tenant) error

	CreateAccount(a *account.Account) error
	CreateTenant(t *tenant.Tenant) error
}
