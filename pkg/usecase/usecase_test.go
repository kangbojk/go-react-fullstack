package usecase

import (
	"testing"

	"github.com/kangbojk/go-react-fullstack/pkg/entity/account"
	"github.com/kangbojk/go-react-fullstack/pkg/entity/tenant"
	"github.com/kangbojk/go-react-fullstack/pkg/storage/memory"
	"github.com/stretchr/testify/assert"
)

func Test_GetTenantUsers(t *testing.T) {
	tRepo := in_memory.NewTenantRepoMem()
	tenant_ := tenant.NewFixtureTenant()
	tenant_.Users = 12

	aRepo := in_memory.NewAccountRepoMem()
	account_ := account.NewFixtureUserAccount()

	tRepo.Create(tenant_)
	aRepo.Create(account_)

	srv := NewService(aRepo, tRepo)

	err := srv.AssignTenantToAccount(account_, tenant_)
	assert.Nil(t, err)

	tenant_users, err := srv.GetTenantUsers(account_)
	assert.Nil(t, err)
	assert.Equal(t, tenant_.Users, tenant_users)
}

func Test_AddUsers(t *testing.T) {
	tRepo := in_memory.NewTenantRepoMem()
	tenant_ := tenant.NewFixtureTenant()

	aRepo := in_memory.NewAccountRepoMem()
	user_account := account.NewFixtureUserAccount()
	admin_account := account.NewFixtureAdminAccount()

	// can use AssignTenantToAccount() instead
	user_account.TenantID = tenant_.ID
	admin_account.TenantID = tenant_.ID

	tRepo.Create(tenant_)
	aRepo.Create(user_account)
	aRepo.Create(admin_account)

	srv := NewService(aRepo, tRepo)

	t.Run("User account not authorized", func(t *testing.T) {
		err := srv.AddUser(user_account, 60)
		assert.Equal(t, account.ErrNotAuthorize, err)
	})

	t.Run("Tenant add users", func(t *testing.T) {
		err := srv.AddUser(admin_account, 60)
		assert.Nil(t, err)
		assert.Equal(t, 60, tenant_.Users)
	})

	t.Run("Tenant full", func(t *testing.T) {
		capacity := tenant_.Capacity
		err := srv.AddUser(admin_account, 2*capacity)
		assert.Equal(t, tenant.ErrFull, err)
	})
}

func Test_Upgrade(t *testing.T) {
	tRepo := in_memory.NewTenantRepoMem()
	tenant_ := tenant.NewFixtureTenant()

	aRepo := in_memory.NewAccountRepoMem()
	user_account := account.NewFixtureUserAccount()
	admin_account := account.NewFixtureAdminAccount()

	// can use AssignTenantToAccount() instead
	user_account.TenantID = tenant_.ID
	admin_account.TenantID = tenant_.ID

	tRepo.Create(tenant_)
	aRepo.Create(user_account)
	aRepo.Create(admin_account)

	srv := NewService(aRepo, tRepo)

	t.Run("User account not authorized", func(t *testing.T) {
		err := srv.Upgrade(user_account)
		assert.Equal(t, account.ErrNotAuthorize, err)
	})

	t.Run("Upgrade to 1000 capacity", func(t *testing.T) {
		err := srv.Upgrade(admin_account)
		assert.Nil(t, err)
		assert.Equal(t, 1000, tenant_.Capacity)
	})

	t.Run("Double upgrade error", func(t *testing.T) {
		err := srv.Upgrade(admin_account)
		assert.Equal(t, tenant.ErrAlreadyUpgraded, err)
		assert.Equal(t, 1000, tenant_.Capacity)
	})
}
