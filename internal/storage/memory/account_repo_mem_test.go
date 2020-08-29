package in_memory

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"gravitational_full_stack_challenge/internal/entity/account"
)

func Test_Account_Create(t *testing.T) {
	repo := NewAccountRepoMem()
	acc := account.NewFixtureAccount()
	id, err := repo.Create(acc)

	assert.Nil(t, err)
	assert.Equal(t, acc.ID, id)
	assert.False(t, acc.CreatedAt.IsZero())
}

func Test_Account_Update(t *testing.T) {
	repo := NewAccountRepoMem()
	acc := account.NewFixtureAccount()
	id, err := repo.Create(acc)

	assert.Nil(t, err)
	saved, _ := repo.Get(id)
	saved.Actions = "R"
	assert.Nil(t, repo.Update(saved))

	updated, err := repo.Get(id)
	assert.Nil(t, err)
	assert.Equal(t, "R", updated.Actions)
}

func Test_Account_Delete(t *testing.T) {
	repo := NewAccountRepoMem()
	acc1 := account.NewFixtureAccount()
	acc2 := account.NewFixtureAccount()

	acc2_id, _ := repo.Create(acc2)

	err := repo.Delete(acc1.ID)
	assert.Equal(t, account.ErrInvalid, err)

	err = repo.Delete(acc2_id)
	assert.Nil(t, err)
	_, err = repo.Get(acc2_id)
	assert.Equal(t, account.ErrInvalid, err)
}
