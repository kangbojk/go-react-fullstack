package account

import (
	"time"

	"gravitational_full_stack_challenge/pkg/ID"
)

func NewFixtureUserAccount() *Account {
	return &Account{
		ID:        id.NewID(),
		Email:     "user11@cheese.com",
		Password:  "manchego",
		Actions:   "R",
		CreatedAt: time.Now(),
	}
}

func NewFixtureAdminAccount() *Account {
	return &Account{
		ID:        id.NewID(),
		Email:     "admin1@cheese.com",
		Password:  "SmokedGouda",
		Actions:   "RW",
		CreatedAt: time.Now(),
	}
}
